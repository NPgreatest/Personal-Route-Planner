<template>
  <div>
    <div class="background2">
      <img :src="imgSrc2" width="100%" height="100%" alt="" />
      <bubbles-effect :options="options" class="bubble"></bubbles-effect>

    </div>



  </div>


</template>

<script>
//import bg1 from "../assets/bg1.jpg"
//import bg2 from "../assets/bg2.jpg"

//import "animate.css"


export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Home",
  data(){
    return{
      imgSrc2:require('../static/bg1.jpg'),
      options: {
        color: 'random', //Bubble color
        radius: 20, //Bubble radius
        densety: 0.4, // The larger the bubble density, the greater the density (suggest no more than 1).
        clearOffset: 0.4 // The larger the bubble disappears [0-1], the longer it disappears.
      },

    }
  }

}
</script>

<style lang="less" scoped>
.background2{
  top:0;
  width:100%;
  height:100%;
  z-index:-1;
  position: absolute;
}
.bubble {
  width: 100%;
}
.animate__animated{
  animation-duration: 3s !important;
}
@keyframes bounce-in {
  0% { transform: translateY(0); }
  20% { transform: translateY(0); }
  50% { transform: translateY(-20px);  border-right: #81f5ac; border-top: #1e703c;}
  80% { transform: translateY(0); }
  100% { transform: translateY(0); }
}


</style>
