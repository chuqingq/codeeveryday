# importing time and vlc
import time, vlc
 
# method to play video
def video(source):
     
    # creating a vlc instance
    vlc_instance = vlc.get_default_instance()
    # vlc_instance.add_intf('')
     
    # creating a media player
    player = vlc_instance.media_player_new()
     
    # creating a media
    media = vlc_instance.media_new(source)

    # setting media to the player
    player.set_media(media)

    inst = player.get_instance()
    # inst.add_intf('')
     
    # play the video
    player.play()
     
    # wait time
    time.sleep(0.5)
     
    # getting the duration of the video
    duration = player.get_length()
     
    # printing the duration of the video
    print("Duration : " + str(duration))
    time.sleep(10)
     
# call the video method
video("1.mp4")