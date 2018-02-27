package http.server;

import static io.netty.handler.codec.http.HttpResponseStatus.OK;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

import io.netty.buffer.ByteBuf;
import io.netty.buffer.Unpooled;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInboundHandlerAdapter;
import io.netty.handler.codec.http.DefaultFullHttpResponse;
import io.netty.handler.codec.http.FullHttpResponse;
import io.netty.handler.codec.http.HttpContent;
import io.netty.handler.codec.http.HttpHeaderNames;
import io.netty.handler.codec.http.HttpRequest;
import io.netty.handler.codec.http.HttpVersion;

public class HttpServerInboundHandler extends ChannelInboundHandlerAdapter {
    private static Log log = LogFactory.getLog(HttpServerInboundHandler.class);

    private HttpRequest request;
    
//    static ByteBuf buf = Unpooled.wrappedBuffer("hello world".getBytes());
    byte[] bs = "hello world".getBytes();

    @Override
    public void channelRead(ChannelHandlerContext ctx, Object msg) throws Exception {
//        if (msg instanceof HttpRequest) {
//            request = (HttpRequest) msg;
//
//            String uri = request.uri();
//            // System.out.println("Uri:" + uri);
//        }
//        if (msg instanceof HttpContent) {
//            HttpContent content = (HttpContent) msg;
//            ByteBuf buf = content.content();
//            // System.out.println(buf.toString(io.netty.util.CharsetUtil.UTF_8));
//            buf.release();

//            String res = "hello world.";
            FullHttpResponse response = new DefaultFullHttpResponse(HttpVersion.HTTP_1_1, OK, Unpooled.wrappedBuffer(bs));
//            response.headers().set(HttpHeaderNames.CONTENT_TYPE, "text/plain");
            response.headers().set(HttpHeaderNames.CONTENT_LENGTH, response.content().readableBytes());
            /*
             * if (HttpHeaders.isKeepAlive(request)) { response.headers().set(CONNECTION, Values.KEEP_ALIVE); }
             */
            ctx.write(response);
//            ctx.flush();
//        }
    }

    @Override
    public void channelReadComplete(ChannelHandlerContext ctx) throws Exception {
        ctx.flush();
    }

    @Override
    public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) {
        log.error(cause.getMessage());
        ctx.close();
    }

}