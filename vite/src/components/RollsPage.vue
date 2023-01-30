<template>
  <div id="rollsPage">
    <h2 class="TitleRollsT">{{ title }}</h2>
    <div class="topic" v-for="(item,index) in rolls.quest" :key="index">
      <span>{{ index + 1 }}.  </span>
      <input class="questTitle" :placeholder="$t('makeQuestBlankTitleNormal')" type="text"
             v-model="item.title">
      <input v-if="item.type === 'blank' || item.type === 'manyBlank'" class="questPlaceholder"
             v-model="item.placeholder" :placeholder="$t('makeQuestTips')">
      <div v-if="item.type === 'choice' || item.type === 'radio'" class="options">
        <input v-for="(options,i) in item.optionsNumber" :key="options + i"
               :placeholder="$t('makeOptions')+(i+1)" class="option" v-model="item.options[i]" type="text">
        <!--               type="text" @input="changeQuestValue(item.options,i,$event.target.value)">-->
      </div>
      <button class="delButton" :title="$t('makeDelete')" @click="deleteQuest(index)">
        <svg fill="red" style="width: 1em;height: 1em;" viewBox="0 0 1025 1024">
          <path
              d="M718.882684 511.351282 1010.410118 800.927611C1024.685139 815.107052 1027.762671 835.153596 1017.28356 845.70286L849.127929 1014.985224C838.649124 1025.534489 818.582055 1022.591441 804.307034 1008.412L512.781438 718.837509 223.267604 1010.296628C209.089387 1024.570118 189.044682 1027.647038 178.496643 1017.169458L9.230209 849.03282C-1.318137 838.554934 1.624604 818.49001 15.802821 804.216827L305.314511 512.759546 13.70467 223.101423C-0.570351 208.921981-3.647577 188.875438 6.831534 178.326173L174.986859 9.043809C185.46597-1.505455 205.533039 1.437592 219.80806 15.617034L511.415756 305.273319 801.039568 13.703302C815.217785-0.570187 835.26249-3.647107 845.810529 6.83078L1015.076963 174.967417C1025.625309 185.444997 1022.682568 205.509921 1008.504351 219.78341L718.882684 511.351282Z"></path>
        </svg>
      </button>
      <!-- v-if来实现各种题目的效果 -->
    </div>
    <div class="line"></div>
  </div>
</template>
<script>
export default {
  name:"RollsPage",
  data(){
    return{
      rolls:{
        title: this.title,
        quest:[
          {
            type: "radio",//或者multipleChoice(多选) / blank(单行填空) manyBlank(多行填空)
            //如果是radio，或者choice，要填写选项数量
            optionsNumber: 3,//选项数量，如果是blank就不用加
            // placeholder: 就是选项相关的提示(placeholder) 如果是选择题就不用加
            title: this.$t("makeQuestTitleNormal"),//题目的标题
            options: [ // 填空题不用加
              this.$t("makeOptions") + "1",
              this.$t("makeOptions") + "2",
              this.$t("makeOptions") + "3"
            ]
          }
        ],
      },
    }
  },
  props:{
    title:String
  },
  methods: {
    changeQuestValue(list, index, value) {
      // this.$set(list, index, value);
      list[index] = value;
    },
    getRoll() {
      return JSON.stringify(this.rolls);
    },
    deleteQuest(index) {
      // delete the quest
      // this.$delete(this.rolls.quest, index)
      // this.rolls.quest.splice(index,1)
      this.rolls.quest.splice(index,1)
    },
    addQuestValue(type, optionsNumber) {
      let quest;
      if (type === "radio" || type === "choice") {
        // if type = choose
        quest = {
          type,
          optionsNumber,
          title: "",
          options: []
        }
        // 用for把数据给怼进去
        for (let i = 0; i < optionsNumber; i++) {
          quest.options.push(this.$t("makeOptions") + (i + 1))
        }
      } else if (type === "blank" || type === "manyBlank") {
        quest = {
          type,
          title: "",
          placeholder: ""
        }
        //This quest not has array, not need for
      }
      this.changeQuestValue(this.rolls.quest, this.rolls.quest.length, quest)
    },
    saveQuest() {
      this.$cookies.config('99999d');
      if (this.$cookies.isKey("draft")) {
        let oldData = this.$cookies.get("draft");
        oldData.data[oldData.data.length] = this.rolls;
        try{
          this.$cookies.set('draft', JSON.stringify(oldData));
        }catch (err){
          return [false,err]
        }
        return [true,null]
        //get after this
      } else {
        try{
          this.$cookies.set('draft', JSON.stringify({data: [this.rolls]}));
        }catch (err){
          return [false,err]
        }
        return [true,null]
      }
    }
  }
}
</script>
<style>
.delButton {
  border: none;
  background: none;
  height: auto;
  position: absolute;
  right: 0;
  bottom: 5px;
}

.options {
  margin-top: 10px;
  margin-left: 10%;
}

.option {
  font-size: 16px;
  border: none;
  outline: none;
  border-bottom: black 1px solid;
  height: 18px;
  margin-top: 10px;
  width: 50%;
  display: block;
}

.option:focus {
  border-bottom: #3b7df0 1px solid;
}

.questTitle {
  font-size: 16px;
  border: none;
  outline: none;
  border-bottom: black 1px solid;
  height: 18px;
  /*margin-top: 10px;*/
  width: calc(100% - 25px);
  /*display: block;*/
}

.line {
  width: 100%;
  height: 1px;
  background: #bebebe;
}

.questTitle:focus {
  border-bottom: #3b7df0 1px solid;
}

.topic {
  width: calc(100% - 60px);

  padding: 10px 30px 10px 30px;
  /*border-bottom: 1px solid #bebebe;*/
  position: relative;
  border-top: 1px solid #bebebe;
  text-align: left;
}

#rollsPage {
  width: calc(50% - 100px);
  box-shadow: 0 3px 10px rgba(0, 0, 0, 0.61);
  padding: 30px 50px 30px 50px;
  margin-top: 100px;
  margin-left: calc(30% + 10%);
  min-height: calc(100% - 123px - 60px);
  border-radius: 1em;
}

.TitleRollsT {
  text-align: center;
}

.questPlaceholder {
  font-size: 16px;
  border: none;
  outline: none;
  border-bottom: black 1px solid;
  height: 18px;
  margin-top: 15px;
  width: 100%;
}
</style>