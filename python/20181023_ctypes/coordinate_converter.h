#ifndef __COORDINATE_CONVERTER_H
#define __COORDINATE_CONVERTER_H

#ifdef __cplusplus
extern "C" {
#endif

void *coordinate_converter_new(double lon, double lat, double alt);
void coordinate_convert(void *converter, double lon0, double lat0, double alt0, double *lon, double *lat, double *alt);
void coordinate_converter_free(void *converter);

#ifdef __cplusplus
}
#endif

#endif