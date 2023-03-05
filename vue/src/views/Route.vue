<template>
  <div class="amap-page-container">
    <div :style="{width:'100%',height:'1080px'}">
      <el-amap vid="amap" :plugin="plugin" :events="events" class="amap-demo" :center="center" >
      </el-amap>
    </div>

  </div>
</template>

<script>
import { lazyAMapApiLoaderInstance } from "vue-amap";
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Route",
  data() {
    return {
      center: [121.59996, 31.197646],
      lng: 0,
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

        const driving = new AMap.Driving({
          map: this.map,
          panel: "panel"
        });
        driving.on("complete", (result) => {
          console.log(result.routes);
        });
        driving.search([
          {keyword: '华东理工大学',city:'上海'},
          {keyword: '龙华烈士陵园',city:'上海'},
        ]);


        // this.drivingMap();
      });
    },
    drivingMap() {
      AMap.plugin("AMap.Driving", () => {
        var driving = new AMap.Driving({
          map: this.map,
          panel: "panel"
        });

        // 根据起终点名称规划驾车导航路线
        driving.search([
          {keyword: '华东理工大学',city:'上海'},
          {keyword: '龙华烈士陵园',city:'上海'},
          {keyword: '中共一大纪念馆',city:'上海'},
          {keyword: '鲁迅故居',city:'上海'}
        ], function(status, result) {
        });
        console.log(driving)
      });

    },
  }
}
</script>

<style scoped>

</style>
