import struct
import sys
import time
from concurrent import futures
from queue import Empty
import grpc
from lib import MiBand3, UUIDS, QUEUE_TYPES

# Generated classes
import mibandDevice_pb2 as mi_pb2
import mibandDevice_pb2_grpc as mi_pb2_grpc

# State of device
isDeviceActive = False
# Notification request interval
requestInterval = 12
# Limit the time to fetch heart beats for
realtimeLimit = 40
# Kill threads with this flag
killProgram = False
# gRPC global state
server = grpc.server(futures.ThreadPoolExecutor(max_workers=1))

def start():
    global server
    mi_pb2_grpc.add_MibandDeviceServicer_to_server(MibandDeviceServicer(), server)
    server.add_insecure_port('[::]:7002')
    server.start()
    print("gRPC server running on port :7002")
    input("Press enter to stop the server")


class MibandDeviceServicer(mi_pb2_grpc.MibandDeviceServicer):

    def GetHeartBeats(self, request, context):
        global isDeviceActive
        global realtimeLimit
        print("Request received from client: ", request.UUID)
        if not isDeviceActive:
            mi = Mi(request.UUID)
            mi.authDevice()
            char_ctrl = mi.setupHeartBeatCalc()

            t = time.time()
            limit = time.time() + realtimeLimit
            while time.time() < limit:
                response = mi_pb2.HeartBeats()
                mi.band.waitForNotifications(0.5)
                hb = mi._parse_queue()
                if (time.time() - t) >= requestInterval:
                    char_ctrl.write(b'\x16', True)
                    t = time.time()
                if hb != -1:
                    print("Streamed heartBeat: ", hb)
                    response.pulse = str(hb)
                    yield response
                else:
                    continue

            print("Done")
            mi.band.stop_realtime()
            mi.band.disconnect()
            isDeviceActive = False
            self.Shutdown()
        
        else:
            response = mi_pb2.HeartBeats()
            response.error = "Device busy"
            yield response
    
    def Shutdown(self):
        server.stop(None)
        print("Killing server")
        sys.exit(0)
    

class Mi:

    def __init__(self, MAC_ADDR):
        global isDeviceActive
        self.band = MiBand3(MAC_ADDR, debug=True)
        isDeviceActive = True

    def authDevice(self):
        print("Device not active")
        self.band.setSecurityLevel(level="medium")
        try:
            self.band.authenticate()
        except:
            self.band.initialize()
            self.band.authenticate()

    def _parse_queue(self):
        while True:
            try:
                res = self.band.queue.get(False)
                _type = res[0]
                if _type == QUEUE_TYPES.HEART:
                    return struct.unpack('bb', res[1])[1]
            except Empty:
                # To keep it seperate from the 0 we get from device
                return -1

    def setupHeartBeatCalc(self):
        char_m = self.band.svc_heart.getCharacteristics(
            UUIDS.CHARACTERISTIC_HEART_RATE_MEASURE)[0]
        char_d = char_m.getDescriptors(
            forUUID=UUIDS.NOTIFICATION_DESCRIPTOR)[0]
        char_ctrl = self.band.svc_heart.getCharacteristics(
            UUIDS.CHARACTERISTIC_HEART_RATE_CONTROL)[0]

        char_sensor = self.band.svc_1.getCharacteristics(
            UUIDS.CHARACTERISTIC_SENSOR)[0]

        # stop heart monitor continues & manual
        char_ctrl.write(b'\x15\x02\x00', True)
        char_ctrl.write(b'\x15\x01\x00', True)
        # WTF
        # char_sens_d1.write(b'\x01\x00', True)
        # enabling accelerometer & heart monitor raw data notifications
        char_sensor.write(b'\x01\x03\x19')
        # IMO: enablee heart monitor notifications
        char_d.write(b'\x01\x00', True)
        # start hear monitor continues
        char_ctrl.write(b'\x15\x01\x01', True)
        # WTF
        char_sensor.write(b'\x02')

        return char_ctrl
