/*
# gcc -I /home/panshangbin/semantic/include sample.c -L /home/panshangbin/semantic/libs -lmsc -ldl -lpthread
# export LD_LIBRARY_PATH=`pwd`/libs
# ./a.out
MSPSearch Succeed, return:
{"rc":0,"operation":"ANSWER","service":"baike","answer":{"text":"周杰伦（Jay Chou），1979年1月18日出生于台湾新北市，华语流行男歌手、词曲创作人、演员、MV及电影导演、编剧及制作人。2000年发行首张个人专辑《Jay》。2002年在中国、新加坡、马来西亚、美国等地举办首场个人世界巡回演唱会。2003年登上美国《时代周刊》亚洲版封面人物。,周杰伦的音乐融合中西方元素，风格多变，四次获得世界音乐大奖“中国区最畅销艺人”奖[3-4]。","type":"T"},"text":"周杰伦"}
enter any key to exit...
*/
#include "stdio.h"
#include "string.h"

#include "msp_cmn.h"
#include "msp_errors.h"

int main(int argc, char** args) {
    const char* login_params			=	"appid = 56d4eb1c, work_dir = .";
    int ret = MSPLogin(NULL, NULL, login_params);
    if (MSP_SUCCESS != ret)
    {
        printf("MSPLogin failed, error code: %d.\n",ret);
	return 0;
    }
    
    const char* params = "nlp_version=2.0";
    const char* text = "周杰伦";
    unsigned int str_len;
    int err;
    const char *result = MSPSearch(params, text, &str_len, &err);
    printf("MSPSearch Succeed, return:\n%s\n", result);
    printf("enter any key to exit...");
    getchar();
    MSPLogout();

    return 0;
}
