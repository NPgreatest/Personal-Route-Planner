<template>

  <div>

    <div class="background2">
      <img :src="imgSrc2" width="100%" height="100%" alt="" />
      <bubbles-effect :options="options" class="bubble"></bubbles-effect>
      <div id="title">

        <div>
          <transition-group appear name="animate__animated animate__bounce animate__slow"
                            enter-active-class="animate__bounceIn"
                            leave-active-class="animate__bounceOut">
            <div v-for="(image, index) in images" :key="index">
              <img :src="image.src" :alt="image.alt" width="150%" height="150%">
            </div>
          </transition-group>
        </div>
<!--        <transition-group appear-->
<!--                          name="animate__animated animate__bounce animate__slow"-->
<!--                          enter-active-class="animate__bounceIn"-->
<!--                          leave-active-class="animate__bounceOut"-->
<!--        >-->
<!--          <div>-->
<!--            <h1>My Image</h1>-->
<!--            <img src="../static/final-logo.png" alt="My Image">-->
<!--          </div>-->
<!--&lt;!&ndash;          <h1 key="1">{{'智能路线规划及其可视化'}}</h1>&ndash;&gt;-->
<!--&lt;!&ndash;          <p key="2">{{'Intelligent-route-planning-and-visualization'}}</p>&ndash;&gt;-->

<!--        </transition-group>-->
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
      imageUrl: '../static/final-logo.png',
      options: {
        color: 'random', //Bubble co
        radius: 20, //Bubble radius1231
        densety: 0.4, // The larger the bubble density, the greater the density (suggest no more than 1).
        clearOffset: 0.4 // The larger the bubble disappears [0-1], the longer it disappears.
      },
      images: [
        { src: require('../assets/final-logo.png'), alt: 'Image 1' },
      ]

    }
  },


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
  top: 30%;
  left: 33%;
  h1 {
    padding-bottom: 20px;
  }

  p {
    font-size: 35px;
  }
}

</style>
