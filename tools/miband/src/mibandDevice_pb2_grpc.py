# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import src.mibandDevice_pb2 as mibandDevice__pb2


class MibandDeviceStub(object):
    """MibandDevice service is to connect to the mibands and operate on it
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetHeartBeats = channel.unary_stream(
                '/MibandDevice/GetHeartBeats',
                request_serializer=mibandDevice__pb2.DeviceUUID.SerializeToString,
                response_deserializer=mibandDevice__pb2.HeartBeats.FromString,
                )


class MibandDeviceServicer(object):
    """MibandDevice service is to connect to the mibands and operate on it
    """

    def GetHeartBeats(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MibandDeviceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetHeartBeats': grpc.unary_stream_rpc_method_handler(
                    servicer.GetHeartBeats,
                    request_deserializer=mibandDevice__pb2.DeviceUUID.FromString,
                    response_serializer=mibandDevice__pb2.HeartBeats.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'MibandDevice', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class MibandDevice(object):
    """MibandDevice service is to connect to the mibands and operate on it
    """

    @staticmethod
    def GetHeartBeats(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/MibandDevice/GetHeartBeats',
            mibandDevice__pb2.DeviceUUID.SerializeToString,
            mibandDevice__pb2.HeartBeats.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)
