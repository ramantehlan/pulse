from src.lib import MiBand3

def start(MAC_ADDR):
    band = MiBand3(MAC_ADDR, debug=True)
    band.setSecurityLevel(level = "medium")
    band.initialize()
    band.authenticate()
    heart_beat(band)

def l(x):
    # This should be streamed
    print('Realtime heart BPM:', x)

def heart_beat(band):
    band.start_raw_data_realtime(heart_measure_callback=l)
    input('Press Enter to continue')
