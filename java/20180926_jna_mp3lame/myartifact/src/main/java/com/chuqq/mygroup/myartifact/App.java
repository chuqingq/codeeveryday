package com.chuqq.mygroup.myartifact;

import java.io.*;

import com.sun.jna.Pointer;

/**
 * Hello world!
 *
 */
public class App {
	public static void main(String[] args) throws Exception {
		System.out.println("main()");

		InputStream in = new FileInputStream("7.5depingfangshiduoshao.pcm");
		OutputStream out = new FileOutputStream("1.mp3");

		Pointer gfp = Mp3Encoder.INSTANCE.lame_init();
		int ret = Mp3Encoder.INSTANCE.lame_set_num_channels(gfp, 1);
		if (ret != 0) {
			System.out.println("lame_set_num_channels error");
			return;
		}

		ret = Mp3Encoder.INSTANCE.lame_set_in_samplerate(gfp, 16000);
		if (ret != 0) {
			System.out.println("lame_set_in_samplerate error");
			return;
		}

		ret = Mp3Encoder.INSTANCE.lame_set_brate(gfp, 16);
		if (ret < 0) {
			System.out.println("lame_set_brate error");
			return;
		}

		ret = Mp3Encoder.INSTANCE.lame_init_params(gfp);
		if (ret < 0) {
			System.out.println("lame_init_params error");
			return;
		}

		final int INBUFSIZE = 4096;
		short[] inputShorts = new short[2 * INBUFSIZE];

		byte[] inputBytes = new byte[2 * inputShorts.length];
		int inputLen = 0;

		final int MP3BUFSIZE = (int) (1.25 * INBUFSIZE) + 7200;
		byte[] outputBytes = new byte[MP3BUFSIZE];
		int outputLen = 0;

		while (true) {
			// 读入inputBytes
			inputLen = in.read(inputBytes, 0, 2 * INBUFSIZE);
			if (inputLen == 0) {
				System.out.println("input finish");
				break;
			}

			if (inputLen < 0) {
				System.out.println("input error");
				break;
			}
			// System.out.println("inputLen: " + inputLen);

			// byte[]转成short[]
			short[] shorts = toShortArray(inputBytes, inputLen);
			System.out.println("input_samples: " + shorts.length + ", first: " + shorts[0]);

			// 转换
			outputLen = Mp3Encoder.INSTANCE.lame_encode_buffer(gfp, shorts, null, shorts.length, outputBytes,
					outputBytes.length);
			if (outputLen < 0) {
				System.out.println("input error");
				break;
			}

			// 拿到outputBytes
			if (outputLen < 0) {
				System.out.println("lame_encode_buffer error");
				break;
			} else if (outputLen > 0) {
				System.out.println("outputLen: " + outputLen);
				// 写文件
				out.write(outputBytes, 0, outputLen);
			}
		}

		if (in != null) {
			in.close();
		}

		if (out != null) {
			out.close();
		}

		if (gfp != Pointer.NULL) {
			Mp3Encoder.INSTANCE.lame_close(gfp);
		}

		System.out.println("finish main()");
	}

	public static short[] toShortArray(byte[] src, int len) {
		int count = len >> 1;
		short[] dest = new short[count];
		for (int i = 0; i < count; i++) {
			dest[i] = (short) (src[i * 2 + 1] << 8 | src[i * 2] & 0xff);
		}
		return dest;
	}

	// public static byte[] toByteArray(short[] src, int offset, int len) {
	// int count = len;
	// byte[] dest = new byte[count << 1];
	// for (int i = offset; i < count; i++) {
	// dest[i * 2 + 1] = (byte) (src[i] >> 8);
	// dest[i * 2] = (byte) (src[i] >> 0);
	// }

	// return dest;
	// }
}
