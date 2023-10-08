<template>
  <!--  <div class="container">
          注释的内容
        </div> -->
  <div class="iframe-container">
    <iframe
        :src="data.externalUrl"
        width="100%"
        height="800px"
        frameborder="0"
        target="_top"
    ></iframe>
  </div>
</template>

// ts(使用了 setup 语法糖)， reactive 自行引入即可
<script setup>
const data = reactive({
externalUrl: 'https://mrhj.gitee.io/form-generator/#/'   // 这里的地址自行替换成自己的
})

</script>