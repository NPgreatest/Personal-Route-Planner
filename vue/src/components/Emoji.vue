<template>
  <div class="emoji-content">
    <div class="emoji">
      <div class="emoji-wrapper">
        开启历史记录后问答助手将会基于历史记录进行回答
      </div>
    </div>
    <div class="mask" @click="closeEmoji"></div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      emojiList: [

      ],
    };
  },
  methods: {
    sendEmoji(item) {
      this.$emit("sendEmoji", item);
    },
    closeEmoji() {
      this.$emit("closeEmoji");
    }
  },
};
</script>
<style lang="scss" scoped>
.emoji-content {
  .emoji {
    width: 420px;
    height: 50px;
    background-color: rgb(39, 42, 55);
    position: absolute;
    top: -70px;
    left: -10px;
    border-radius: 10px;
    transition: 0.2s;
    z-index: 4;

    &::after {
      content: "";
      display: block;
      width: 0;
      height: 0;
      border-top: 10px solid rgb(39, 42, 55);
      border-right: 10px solid transparent;
      border-left: 10px solid transparent;
      position: absolute;
      bottom: -8px;
      left: 15px;
      z-index: 100;
    }

    .emoji-wrapper {
      color: white;  // Set text color to white
      font-weight: bold;  // Make text bold
      width: 100%;
      height: 100%;
      overflow-y: scroll;
      padding: 10px;
      box-sizing: border-box;
      position: relative;

      &::-webkit-scrollbar {
        /* Scrollbar overall style */
        width: 4px; /* Dimensions correspond to the width and height of the vertical and horizontal scrollbars, respectively */
        height: 1px;
      }
      &::-webkit-scrollbar-thumb {
        /* Scrollbar thumb block */
        border-radius: 10px;
        box-shadow: inset 0 0 5px rgba(97, 184, 179, 0.1);
        background: rgb(95, 101, 122);
      }
      &::-webkit-scrollbar-track {
        /* Scrollbar track */
        box-shadow: inset 0 0 5px rgba(87, 175, 187, 0.1);
        border-radius: 10px;
        background: rgb(39, 42, 55);
      }
    }
  }

  .mask {
    width: 100%;
    height: 100%;
    position: fixed;
    background: transparent;
    left: 0;
    top: 0;
    z-index: 1;
  }
}
</style>
