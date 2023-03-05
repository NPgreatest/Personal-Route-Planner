<template>
  <div class="amap-page-container">
    <div :style="{width:'100%',height:'1080px'}">
      <el-amap vid="amap" :plugin="plugin" class="amap-demo" :center="center">
      </el-amap>
    </div>

  </div>
</template>

<script>
// import "@/plugins/aMap.js";
import { lazyAMapApiLoaderInstance } from "vue-amap";
// import AMap from 'vue-amap';
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Route",
  data(){
    const self = this;
    return {
      center: [121.59996, 31.197646],
      lng: 0,
      lat: 0,
      loaded: false,
      plugin: [{
        enableHighAccuracy: true,//是否使用高精度定位，默认:true
        timeout: 100,          //超过10秒后停止定位，默认：无穷大
        maximumAge: 0,           //定位结果缓存0毫秒，默认：0
        convert: true,           //自动偏移坐标，偏移后的坐标为高德坐标，默认：true
        showButton: true,        //显示定位按钮，默认：true
        buttonPosition: 'RB',    //定位按钮停靠位置，默认：'LB'，左下角
        showMarker: true,        //定位成功后在定位到的位置显示点标记，默认：true
        showCircle: true,        //定位成功后用圆圈表示定位精度范围，默认：true
        panToLocation: true,     //定位成功后将定位到的位置作为地图中心点，默认：true
        zoomToAccuracy:true,//定位成功后调整地图视野范围使定位位置及精度范围视野内可见，默认：f
        extensions:'all',
        pName: 'Geolocation',
        events: {
          init(o) {
            // o 是高德地图定位插件实例
            o.getCurrentPosition((status, result) => {
              console.log(result)
              if (result && result.position) {
                self.lng = result.position.lng;
                self.lat = result.position.lat;
                self.center = [self.lng, self.lat];
                self.loaded = true;
                self.$nextTick();
              }
            });
          }
        }
      }]
    }
  },
  mounted() {
    lazyAMapApiLoaderInstance.load().then(() => {
      this.map = new AMap.Map("amapContainer", {
        center: new AMap.LngLat(114.268691, 30.401227),
        zoom: this.zoom,
      });
      this.drivingMap();
    });
    // this.initMap();
  },
  methods:{
    drivingMap(){
      AMap.service(["AMap.PlaceSearch"], function() {
        //构造地点查询类
        var placeSearch = new AMap.PlaceSearch({
          pageSize: 5, // 单页显示结果条数
          pageIndex: 1, // 页码
          city: "上海市", // 兴趣点城市
          citylimit: true,  //是否强制限制在设置的城市内搜索
          map: map, // 展现结果的地图实例
          panel: "panel1", // 结果列表将在此容器中进行展示。
          autoFitView: true // 是否自动调整地图视野使绘制的 Marker点都处于视口的可见范围
        });
        //关键字查询
        placeSearch.search('{{ destination }}');
      });
      var driving = new AMap.Driving({
        map: map,
        panel: "panel"
      });
      // 根据起终点名称规划驾车导航路线
      driving.search([
        {keyword: '华东理工大学',city:'上海'},
        {keyword: '龙华烈士陵园',city:'上海'},
        {keyword: '中共一大纪念馆',city:'上海'},
        {keyword: '鲁迅故居',city:'上海'}
      ], function(status, result) {
        // result 即是对应的驾车导航信息，相关数据结构文档请参考  https://lbs.amap.com/api/javascript-api/reference/route-search#m_DrivingResult
        if (status === 'complete') {
          log.success('绘制驾车路线完成')
        } else {
          log.error('获取驾车数据失败：' + result)
        }
      });
    },
    initMap() {

      // AMap.load({
      //   key: "73674da9f8fe033be85592e5c836dc4c", //此处填入我们注册账号后获取的Key
      //   //version: "2.0", //指定要加载的 JSAPI 的版本，缺省时默认为 1.4.15
      //   plugins: ['AMap.Weather','AMap.PlaceSearch'], //需要使用的的插件列表，如比例尺'AMap.Scale'等
      // }).then((AMap) => {
      //   this.map = new AMap.Map("container", { //设置地图容器id
      //     viewMode: "3D", //是否为3D地图模式
      //     zoom: 13, //初始化地图级别
      //     center: [121.436509,31.189712],//地图中心点
      //   });
      // }).catch(e => {
      //   console.log(e);
      // })
      AMap.plugin('AMap.Weather', function() {
        var weather = new AMap.Weather();
        //查询实时天气信息, 查询的城市到行政级别的城市，如朝阳区、杭州市
        weather.getLive('上海市', function(err, data) {
          if (!err) {
            var str = [];
            str.push('<h4 >实时天气' + '</h4><hr>');
            str.push('<p>城市/区：' + data.city + '</p>');
            str.push('<p>天气：' + data.weather + '</p>');
            str.push('<p>温度：' + data.temperature + '℃</p>');
            str.push('<p>风向：' + data.windDirection + '</p>');
            str.push('<p>风力：' + data.windPower + ' 级</p>');
            str.push('<p>空气湿度：' + data.humidity + '</p>');
            str.push('<p>发布时间：' + data.reportTime + '</p>');
            var marker = new AMap.Marker({map: map, position: map.getCenter()});
            var infoWin = new AMap.InfoWindow({
              content: '<div class="info" style="position:inherit;margin-bottom:10;" onclick="disappear()" id="w">'+str.join('')+'</div><div class="sharp"></div>',
              isCustom:true,
              offset: new AMap.Pixel(0, -37)
            });
            infoWin.open(map, marker.getPosition());
            marker.on('mouseover', function() {
              infoWin.open(map, marker.getPosition());
            });
          }
        });

      });
      AMap.service(["AMap.PlaceSearch"], function() {
            //构造地点查询类
            var placeSearch = new AMap.PlaceSearch({
              pageSize: 5, // 单页显示结果条数
              pageIndex: 1, // 页码
              city: "上海市", // 兴趣点城市
              citylimit: true,  //是否强制限制在设置的城市内搜索
              map: map, // 展现结果的地图实例
              panel: "panel1", // 结果列表将在此容器中进行展示。
              autoFitView: true // 是否自动调整地图视野使绘制的 Marker点都处于视口的可见范围
            });
            //关键字查询
            placeSearch.search('{{ destination }}');
          });
          var driving = new AMap.Driving({
            map: map,
            panel: "panel"
          });
          // 根据起终点名称规划驾车导航路线
          driving.search([
            {keyword: '华东理工大学',city:'上海'},
            {keyword: '龙华烈士陵园',city:'上海'},
        {keyword: '中共一大纪念馆',city:'上海'},
        {keyword: '鲁迅故居',city:'上海'}
      ], function(status, result) {
        // result 即是对应的驾车导航信息，相关数据结构文档请参考  https://lbs.amap.com/api/javascript-api/reference/route-search#m_DrivingResult
        if (status === 'complete') {
          log.success('绘制驾车路线完成')
        } else {
          log.error('获取驾车数据失败：' + result)
        }
      });




    },
  },
}
</script>

<style scoped>
#container {
  width: 100%;
  height: 850px;
  margin: 100px auto;
  border: 2px solid red;
}
</style>
