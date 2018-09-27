package com.chuqq.mygroup.myartifact;

import com.sun.jna.Library;
import com.sun.jna.Native;
// import com.sun.jna.Structure;
import com.sun.jna.Pointer;

/**
 * Hello world!
 *
 */
public interface Mp3Encoder extends Library {
	public static Mp3Encoder INSTANCE = (Mp3Encoder) Native.loadLibrary("mp3lame", Mp3Encoder.class);

	Pointer lame_init();
	int lame_set_num_channels(Pointer gfp, int num_channels);
	int lame_set_in_samplerate(Pointer gfp, int in_samplerate);
	int lame_set_brate(Pointer gfp, int brate);
	int lame_init_params(Pointer gfp);
	int lame_encode_buffer(Pointer gfp, short[] inputL, short[] inputR, int input_samples, byte[] mp3_buffer, int mp3_buffer_len);
	int lame_close(Pointer gfp);
}
