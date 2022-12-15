<template>
  <span class="headerAvatar">
    <template v-if="picType === 'avatar'">
      <el-avatar v-if="userStore.userInfo.headerImg" :size="30" :src="avatar" />
      <el-avatar v-else :size="30" :src="noAvatar" />
    </template>
    <template v-if="picType === 'img'">
      <img v-if="userStore.userInfo.headerImg" :src="avatar" class="avatar">
      <img v-else :src="noAvatar" class="avatar">
    </template>
    <template v-if="picType === 'file'">
      <img :src="file" class="file">
    </template>

    <template v-if="picType === 'carousel' ">
      <div class="demo-image__preview" style="z-index: 2009;">
        <el-image
          style=" height: 100px"
          :src="imgUrl"
          :preview-src-list="srcList"
          :initial-index="1"
          :hide-on-click-modal="true"
          fit="cover"
          :preview-teleported="true"
        />
      </div>
    </template>
  </span>
</template>

<script>
export default {
  name: 'CustomPic'
}
</script>

<script setup>
import noAvatarPng from '@/assets/noBody.png'
import { useUserStore } from '@/pinia/modules/user'
import { computed, ref } from 'vue'
const props = defineProps({
  picType: {
    type: String,
    required: false,
    default: 'avatar'
  },
  picSrc: {
    type: String,
    required: false,
    default: ''
  }
})


const srcList = [
  
]

const path = ref(import.meta.env.VITE_BASE_API + '/')
const noAvatar = ref(noAvatarPng)

const userStore = useUserStore()

const avatar = computed(() => {
  if (props.picSrc === '') {
    if (userStore.userInfo.headerImg !== '' && userStore.userInfo.headerImg.slice(0, 4) === 'http') {
      return userStore.userInfo.headerImg
    }
    return path.value + userStore.userInfo.headerImg
  } else {
    if (props.picSrc !== '' && props.picSrc.slice(0, 4) === 'http') {
      return props.picSrc
    }
    return path.value + props.picSrc
  }
})
const file = computed(() => {
  if (props.picSrc && props.picSrc.slice(0, 4) !== 'http') {
    return path.value + props.picSrc
  }
  return props.picSrc
})

const imgUrl = computed(() => {
  if (props.picSrc && props.picSrc.slice(0, 4) !== 'http') {
    srcList.push(path.value + props.picSrc)
    return path.value + props.picSrc
  }
  srcList.push(props.picSrc)
  return props.picSrc
})

</script>

<style scoped>
.headerAvatar{
    display: flex;
    justify-content: center;
    align-items: center;
    margin-right: 8px;
}
.demo-image__preview{
  position:absolute;

}
.file{
    /* width: 80px; */
    height: 80px;
    position: relative;
}

.demo-image__error .image-slot {
  font-size: 30px;
}
.demo-image__error .image-slot .el-icon {
  font-size: 30px;
}
.demo-image__error .el-image {
  width: 100%;
  height: 200px;
}

</style>
