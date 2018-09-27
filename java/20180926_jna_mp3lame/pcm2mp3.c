
#include <stdio.h>
#include <stdlib.h>

#include "lame/lame.h"

int encode(const char* inPath, const char* outPath) 
{
	// 打开pcm文件
	FILE* infp = fopen(inPath, "rb");
	if (infp == NULL) {
		printf("inPath %s open error\n", inPath);
		return -1;
	}

	// 打开mp3文件
	FILE* outfp = fopen(outPath, "wb");
	if (outfp == NULL) {
		printf("outPath %s open error\n", outPath);
		fclose(infp);
		return -2;
	}
 
	// 初始化lame参数
	lame_global_flags* gfp;
	gfp = lame_init();
	if (gfp == NULL) 
	{
		printf("lame_init failed/n");
		fclose(infp);
		fclose(outfp);
		return -3;
	}

	// lame参数
	lame_set_num_channels(gfp, 1);
	lame_set_in_samplerate(gfp, 16000);

	int ret = lame_init_params(gfp);
	if (ret < 0) 
	{
		printf("lame_init_params returned %d\n",ret);
		return -4;
	}
 
	// 输入缓冲区
	const int INBUFSIZE = 4096;
	short input_buffer[INBUFSIZE*2];
	int input_samples;

	// 输出缓冲区
	const int MP3BUFSIZE = (int) (1.25 * INBUFSIZE) + 7200;
	unsigned char mp3_buffer[MP3BUFSIZE];
	int mp3_bytes;
 
	while(1) {
		input_samples = fread(input_buffer, 2, INBUFSIZE, infp);
		// printf("input_samples is %d.\n",input_samples);
		if (input_samples == 0) {
			break;
		}
		else if (input_samples < 0) {
			printf("input_samples invalid\n");
			break;
		}
		printf("input_samples: %d, first: %d\n", input_samples, input_buffer[0]);

		// mp3_bytes = lame_encode_buffer_interleaved(gfp, input_buffer, input_samples/2, mp3_buffer, sizeof(mp3_buffer));
		mp3_bytes = lame_encode_buffer(gfp, input_buffer, NULL, input_samples, mp3_buffer, sizeof(mp3_buffer));

		if (mp3_bytes < 0) {
			printf("lame_encode_buffer_interleaved returned %d/n", mp3_bytes);
		}
		else if(mp3_bytes > 0) {
			printf("outputLen: %d\n", mp3_bytes);
			fwrite(mp3_buffer, 1, mp3_bytes, outfp);
		}
	}
 
	if (outfp != NULL)
		fclose(outfp);

	if (infp != NULL)
		fclose(infp);

	if (gfp != NULL)
		lame_close(gfp);

	return 0;
}
 
int main(int argc, char** argv) 
{
	if (argc < 3) 
	{
		printf("usage: pcm2mp3 rawinfile mp3outfile/n");
		return -1;
	}
	encode(argv[1], argv[2]);
	return 0;
}

