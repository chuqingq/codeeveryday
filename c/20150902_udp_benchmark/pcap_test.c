#include <pcap.h>
#include <time.h>
#include <stdlib.h>
#include <stdio.h>

int count = 0;
long long before = 0;

long long msec(void) {
    struct timeval tv;
    gettimeofday(&tv, NULL);
    return (((long long)tv.tv_sec) * 1000) + tv.tv_usec / 1000;
}

void getPacket(u_char * arg, const struct pcap_pkthdr * pkthdr, const u_char * packet)
{
    // int * id = (int *)arg;

    // printf("id: %d\n", ++(*id));
    // printf("Packet length: %d\n", pkthdr->len);
    // printf("Number of bytes: %d\n", pkthdr->caplen);
    // printf("Recieved time: %s", ctime((const time_t *)&pkthdr->ts.tv_sec));

    // int i;
    // for (i = 0; i < pkthdr->len; ++i)
    // {
    //     printf(" %02x", packet[i]);
    //     if ( (i + 1) % 16 == 0 )
    //     {
    //         printf("\n");
    //     }
    // }

    // printf("\n\n");

    count++;
    if (count >= 100000) {
        long long now = msec();
        printf("recv %f\n", count * 1.0 / (now - before));

        count = 0;
        before = now;
    }
}

int main()
{
    char errBuf[PCAP_ERRBUF_SIZE], * devStr;

    /* get a device */
    devStr = pcap_lookupdev(errBuf);

    if (devStr)
    {
        printf("success: device: %s\n", devStr);
    }
    else
    {
        printf("error: %s\n", errBuf);
        exit(1);
    }

    /* open a device, wait until a packet arrives */
    pcap_t * device = pcap_open_live("eth1", 65535, 1, 0, errBuf);

    if (!device)
    {
        printf("error: pcap_open_live(): %s\n", errBuf);
        exit(1);
    }

    before = msec();
    /* wait loop forever */
    int id = 0;
    pcap_loop(device, -1, getPacket, (u_char*)&id);

    pcap_close(device);

    return 0;
}
