<template>
  <div class="summary-container">
    <h1 class="title">团日活动总结</h1>
    <div class="summary" v-html="formattedSummary"></div>
  </div>
</template>

<script>
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Summary",
  data() {
    return {
      summary: ""
    };
  },
  computed: {
    formattedSummary() {
      return this.summary.replace(/\n/g, "<br>");
    }
  },
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.summary = decodeURIComponent( localStorage.getItem("summay"));
    }
  }
};
</script>

<style scoped>
.summary-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f2f2f2;
}

.title {
  font-size: 2.5rem;
  font-weight: bold;
  text-align: center;
  margin-bottom: 30px;
}

.summary {
  width: 90%;
  max-width: 800px;
  background-color: #ffffff;
  border: 1px solid #ccc;
  padding: 30px;
  margin-bottom: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

@media (min-width: 768px) {
  .summary {
    width: 80%;
  }
}
</style>
