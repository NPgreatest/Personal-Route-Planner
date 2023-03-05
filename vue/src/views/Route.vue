<template>
  <div class="amap-page-container">
    <div :style="{width:'100%',height:'1080px'}">
      <el-amap vid="amap" :plugin="plugin"  :events="events" class="amap-demo" :center="center" >
        <div id="panel" />
      </el-amap>
    </div>

  </div>
</template>

<script>

import { lazyAMapApiLoaderInstance } from "vue-amap";
export default {

  // eslint-disable-next-line vue/multi-word-component-names
  name: "Route",
  props:['routes'],
  data() {
    return {
      center: [121.59996, 31.197646],
      lng: 0,
      path:[],
      events:{},
      lat: 0,
      loaded: false,
      plugin: ["AMap.Driving"],
    }
  },
  mounted() {
    this.initMap();
  },
  methods: {
    initMap() {

      lazyAMapApiLoaderInstance.load().then(() => {
        this.map = new AMap.Map("amap", {
          center: [121.59996, 31.197646],
          zoom: this.zoom,
        });
         this.drivingMap();
      });
    },
    drivingMap() {
      AMap.plugin("AMap.Driving", () => {
        var driving = new AMap.Driving({
          map: this.map,
          panel: "panel"
        });
        const routes =JSON.parse(localStorage.getItem('routes'))
        // 根据起终点名称规划驾车导航路线
        driving.search([
          {keyword: routes[0].toString() ,city:'上海'},
          {keyword: routes[1].toString() ,city:'上海'},
          {keyword: routes[2].toString() ,city:'上海'},
          {keyword: routes[3].toString() ,city:'上海'}
        ], function(status, result) {
          console.log(status,result)
          if (status === 'complete') {
            self.path = result.routes[0].path;
            self.renderRoute();
          }
        });
        console.log(driving)
      });
      const path = this.path;
      if (path.length > 0) {
        const polyline = new AMap.Polyline({
          map: this.map,
          path: path,
          strokeColor: '#3366FF',
          strokeOpacity: 0.8,
          strokeWeight: 5,
          strokeStyle: 'solid'
        });
        this.map.setFitView(polyline);
      }
    }
  }
}
</script>

<style scoped>

</style>
