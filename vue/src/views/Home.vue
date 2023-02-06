<template>
  <div>
    <div class="background2">
      <img :src="imgSrc2" width="100%" height="100%" alt="" />
      <bubbles-effect :options="options" class="bubble"></bubbles-effect>
      <div id="title">
        <transition-group appear
                          name="animate__animated animate__bounce animate__slow"
                          enter-active-class="animate__bounceIn"
                          leave-active-class="animate__bounceOut"
        >
          <h1 key="1">{{'智能路线规划及其可视化'}}</h1>
          <p key="2">{{'Intelligent-route-planning-and-visualization'}}</p>
        </transition-group>
      </div>
    </div>



  </div>


</template>

<script>
//import bg1 from "../assets/bg1.jpg"
//import bg2 from "../assets/bg2.jpg"

import "animate.css"


export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Home",
  data(){
    return{
      imgSrc2:require('../static/background.jpg'),
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
  //top:0;
  width:100%;
  height:100%;
  z-index:-1;
  position: absolute;
  background-size: 100% 100%;
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
#title {
  //z-index:1;
  position: absolute;
  user-select: none;
  font-weight: 760;
  line-height: 1.25;
  color: white;
  text-align: center;
  top: 42%;
  left: 30%;
  h1 {
    padding-bottom: 20px;
  }

  p {
    font-size: 35px;
  }
}

</style>
