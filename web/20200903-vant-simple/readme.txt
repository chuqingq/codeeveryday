遇到calendar无法关闭的问题，把标签改成带闭合的，就可以了。即：
<van-calendar v-model="show" @confirm="onConfirm" />
改成
<van-calendar v-model="show" @confirm="onConfirm" ></van-calendar>
