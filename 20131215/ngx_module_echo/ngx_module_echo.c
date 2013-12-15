#include <ngx_config.h>
#include <ngx_core.h>
#include <ngx_http.h>

static u_char ngx_hello_world[]="hello world";

static ngx_int_t
ngx_echo_handler(ngx_http_request_t *r)
{
    printf("called:ngx_echo_handler\n");
    ngx_buf_t *b; 
    ngx_chain_t out; 

    r->headers_out.content_type.len =sizeof("text/plain")-1; 
    r->headers_out.content_type.data=(u_char *) "text/plain"; 

    b=ngx_pcalloc(r->pool,sizeof(ngx_buf_t)); 

    out.buf=b;  
    out.next=NULL; 

    b->pos=ngx_hello_world;  
    b->last=ngx_hello_world+sizeof(ngx_hello_world); 
    b->memory=1; 
    b->last_buf=1; 

    r->headers_out.status=NGX_HTTP_OK; 
    r->headers_out.content_length_n=sizeof(ngx_hello_world); 
    ngx_http_send_header(r); 

    return ngx_http_output_filter(r, &out);
}

static char *
ngx_echo_readconf(ngx_conf_t *cf, ngx_command_t *cmd, void *conf)
{
    printf("called:ngx_echo_readconf\n");
    ngx_http_core_loc_conf_t  *clcf;

    clcf = ngx_http_conf_get_module_loc_conf(cf, ngx_http_core_module);
    clcf->handler = ngx_echo_handler;
    // ngx_conf_set_str_slot(cf,cmd,conf);
    return NGX_CONF_OK;
}

static ngx_command_t  ngx_echo_commands[] = {
    { ngx_string("echo"),
      NGX_HTTP_LOC_CONF|NGX_CONF_TAKE1,
      ngx_echo_readconf,
      NGX_HTTP_LOC_CONF_OFFSET,
      0,
      NULL },
      ngx_null_command
};

static ngx_http_module_t  ngx_echo_module_ctx = {
    NULL, /* preconfiguration */
    NULL, /* postconfiguration */

    NULL, /* create main configuration */
    NULL, /* init main configuration */

    NULL, /* create server configuration */
    NULL, /* merge server configuration */

    NULL, /* create location configuration */
    NULL, /* merge location configuration */
};

ngx_module_t  ngx_module_echo = {
    NGX_MODULE_V1,
    &ngx_echo_module_ctx,          /* module context */
    ngx_echo_commands,             /* module directives */
    NGX_HTTP_MODULE,               /* module type */
    NULL,                          /* init master */
    NULL,                          /* init module */
    NULL,                          /* init process */
    NULL,                          /* init thread */
    NULL,                          /* exit thread */
    NULL,                          /* exit process */
    NULL,                          /* exit master */
    NGX_MODULE_V1_PADDING
};
