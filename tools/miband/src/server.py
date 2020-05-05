import time
from src.lib import MiBand3

# for 4 minutes
realtimeLimit = time.time() + 60 * 4


class start:

    def __init__(self, MAC_ADDR):
        self.band = MiBand3(MAC_ADDR, debug=True)
        # Start the server here
        self.runDevice(MAC_ADDR)


    def runDevice(self, MAC_ADDR):
        self.band.setSecurityLevel(level="medium")
        self.band.initialize()
        self.band.authenticate()
        self.startHeartBeatCalc(self.band)


    def streamHeartBeat(self, x):
        # This should be streamed
        print('Realtime heart BPM:', x)
        # This should not be a if, but in while loop where the func is defined
        # But for some reason it's not working like that in legacy code, so
        # this is just temp fix
        if time.time() > realtimeLimit:
            print('stop')
            self.band.stop_realtime()


    def startHeartBeatCalc(self, band):
        self.band.start_raw_data_realtime(
                heart_measure_callback=self.streamHeartBeat
        )
