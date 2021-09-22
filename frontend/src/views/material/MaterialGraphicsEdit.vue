<template>
  <v-app>
    <v-row>
      <v-card width="660px" class="mx-auto my-12">
        <v-card-title>
          <v-row>
            <v-col cols="12"
                   lg="9"
                   md="8"
                   sm="6"
                   xs="12">
              <v-text-field id="title" v-model="title" outlined dense label="标题"></v-text-field>
            </v-col>
            <v-col cols="12"
                   lg="3"
                   md="4"
                   sm="6"
                   xs="12">
              <v-select
                  v-model="channel"
                  item-text="name"
                  item-value="id"
                  :items="channelItems"
                  label="所属频道"
                  dense
                  outlined
              ></v-select>
            </v-col>
          </v-row>
        </v-card-title>
        <v-card-text>
          <v-row
              align="center"
              class="mx-0 tinymce-content"
          >
            <tinymce-message-editor @hook:destroyed="editorInit" v-if="showEditor" :editor-templates="templates" :channel="channel" base-url="tinymce/"
                                    v-model="content"></tinymce-message-editor>
          </v-row>
        </v-card-text>
        <v-divider class="mx-4"></v-divider>
        <v-card-subtitle>
          <v-row>

            <div class="text-h4 pa-3">封面和摘要</div>

            <v-spacer></v-spacer>
          </v-row>

          <v-row>
            <v-col>
              <cover-image v-on:getCoverImageUrl="getCoverImageUrl" :channel="channel"></cover-image>
            </v-col>
            <v-divider vertical></v-divider>
            <v-col>
              <v-textarea label="摘要" v-model="abstract" filled
                          rows="4"
                          row-height="30"
              ></v-textarea>
            </v-col>
          </v-row>
        </v-card-subtitle>
        <v-divider class="mx-4"></v-divider>
        <v-card-actions class="pa-3">
          <v-spacer></v-spacer>
          <v-btn
              color="success"
              @click="store"
          >
            保存
          </v-btn>
          <v-btn
              color="text"
              @click="preview"
          >
            预览
          </v-btn>
          <v-btn
              color="text"
              @click="publish"
          >
            保存并群发
          </v-btn>
        </v-card-actions>
      </v-card>
      <v-overlay :value="overlay">
        <v-progress-circular
            indeterminate
            color="primary"
        ></v-progress-circular>
      </v-overlay>
    </v-row>
    <v-dialog v-model="dialog" width="960">
      <v-card>
        <v-card-title>
          <h2>群发消息</h2>
        </v-card-title>
        <v-card-text class="pa-4">
          <v-card>
            <v-card-title>
              <v-row class="pa-4">
                <h4>选择用户</h4>
                <v-tooltip bottom>
                  <template v-slot:activator="{ on, attrs }">
                    <v-icon
                        v-bind="attrs"
                        v-on="on"
                        color="blue darken-2"
                    >
                      mdi-help-circle-outline
                    </v-icon>
                  </template>
                  <span>选择部门后，该部门下的所有员工都会接收到该消息</span>
                </v-tooltip>
              </v-row>
            </v-card-title>
            <v-card-text>
              <v-row>
                <v-col>
                  <v-text-field
                      v-model="search"
                      label="输入关键字搜索"
                      clearable
                      @input="handleSearch"
                      clear-icon="mdi-close-circle-outline"
                  ></v-text-field>
                  <v-treeview
                      ref="items"
                      v-model="selection"
                      :items="items"
                      selectable
                      return-object
                      dense
                      hoverable
                      :open.sync="open"
                      item-text="text"
                      :search="search"
                  ></v-treeview>
                </v-col>
                <v-divider vertical></v-divider>
                <v-col>
                  <v-card-text>
                    <v-scroll-x-transition
                        group
                        hide-on-leave
                    >
                      <v-chip
                          v-for="(node, i) in selection"
                          :key="i"
                          color="primary"
                          outlined
                          small
                          close
                          class="ma-1"
                          @click:close="selection.splice(i,1)"
                      >
                        <v-icon
                            left
                            small
                        >
                          mdi-account
                        </v-icon>
                        {{ node.text }}
                      </v-chip>
                    </v-scroll-x-transition>
                  </v-card-text>
                </v-col>
              </v-row>
            </v-card-text>
            <v-card-subtitle>
              <v-container fluid>
                <v-row align="center" justify="center">
                  <v-col cols="12" md="3" xs="12">
                    <v-switch
                        v-model="timed_sending"
                        label="定时群发"
                    ></v-switch>
                  </v-col>
                  <v-col v-if="timed_sending" cols="12" md="3" xs="12">
                    <date-picker v-model="send_time" :disabled-date="notBeforeToday" type="datetime" value-type="format"
                                 format="YYYY-MM-DD HH:mm:ss"></date-picker>
                  </v-col>
                  <v-spacer></v-spacer>
                </v-row>
                <v-row align="center" justify="center">
                  <v-col cols="12" md="3" xs="12">
                    <v-switch
                        v-model="author_switch"
                        label="是否填写作者"
                    ></v-switch>
                  </v-col>
                  <v-col v-if="author_switch" cols="12" md="3" xs="12">
                    <v-text-field v-model="author" outlined dense label="作者"></v-text-field>
                  </v-col>
                  <v-spacer></v-spacer>
                </v-row>
              </v-container>
            </v-card-subtitle>
          </v-card>
        </v-card-text>
        <v-card-actions>
          <v-btn
              color="success"
              @click="store(0)"
              class="text-center"
          >
            群发
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-app>
</template>

<script>
let moment = require("moment");
import CoverImage from "../../components/material/image/CoverImage";
import TinymceMessageEditor from "../../components/tinymce/TinymceMessageEditor";
import DatePicker from 'vue2-datepicker';
import 'vue2-datepicker/index.css';

export default {
  name: "GraphicsMessageEdit",
  components: {TinymceMessageEditor, CoverImage, DatePicker},
  data() {
    return {
      open: [],
      allOpened: false,
      lastOpen: [],
      overlay: false,
      timed_sending: false,
      author_switch: false,
      author: "",
      send_time: moment().format('YYYY-MM-DD HH:mm:ss'),
      channel: 0,
      channelItems: [{id:1,name:'gy'},{id:2,name:'afgasf'}],
      showEditor: true,
      list: [],
      search: null,
      selection: [],
      items: [{
        "id": "1",
        "type": "department",
        "text": "快牛金科",
        "children": [{
          "id": "302",
          "type": "department",
          "text": "B1-钱牛牛中心",
          "children": [{
            "id": "303",
            "type": "department",
            "text": "钱牛牛运营组",
            "children": [{"id": "307", "type": "department", "text": "运营组"}, {
              "id": "308",
              "type": "department",
              "text": "客服组"
            }, {"id": "306", "type": "department", "text": "推广组"}]
          }, {"id": "304", "type": "department", "text": "钱牛牛产品组"}, {
            "id": "305",
            "type": "department",
            "text": "钱牛牛研发组",
            "children": [{"id": "309", "type": "department", "text": "后端开发组"}, {
              "id": "311",
              "type": "department",
              "text": "移动开发组",
              "children": [{"id": "310", "type": "department", "text": "前端开发组1"}]
            }, {"id": "313", "type": "department", "text": "测试组"}, {"id": "325", "type": "department", "text": "BI组"}]
          }, {"id": "317", "type": "department", "text": "钱牛牛法务组"}, {
            "id": "345",
            "type": "department",
            "text": "钱牛牛财务组"
          }, {"id": "359", "type": "department", "text": "钱牛牛设计组"}, {
            "id": "360",
            "type": "department",
            "text": "钱牛牛品牌组"
          }]
        }, {
          "id": "4",
          "type": "department",
          "text": "管理层",
          "children": [{"id": "388", "type": "department", "text": "管理中层外部2"}]
        }, {
          "id": "193",
          "type": "department",
          "text": "B2-资产业务中心",
          "children": [{"id": "286", "type": "department", "text": "产品组"}, {
            "id": "290",
            "type": "department",
            "text": "运营组"
          }, {
            "id": "287",
            "type": "department",
            "text": "开发组",
            "children": [{"id": "292", "type": "department", "text": "后端开发组"}, {
              "id": "293",
              "type": "department",
              "text": "前端开发组"
            }, {"id": "294", "type": "department", "text": "移动开发组"}, {
              "id": "295",
              "type": "department",
              "text": "测试组"
            }, {"id": "296", "type": "department", "text": "设计组"}]
          }]
        }, {
          "id": "384",
          "type": "department",
          "text": "B3-企业服务中心",
          "children": [{"id": "1002", "type": "department", "text": "商务组"}, {
            "id": "1001",
            "type": "department",
            "text": "产品组"
          }, {
            "id": "218",
            "type": "department",
            "text": "研发组",
            "children": [{"id": "230", "type": "department", "text": "移动开发组"}]
          }]
        }, {
          "id": "186",
          "type": "department",
          "text": "O1-风控中心",
          "children": [{"id": "326", "type": "department", "text": "BD组"}, {
            "id": "327",
            "type": "department",
            "text": "产品组"
          }, {"id": "188", "type": "department", "text": "风控策略组"}, {
            "id": "189",
            "type": "department",
            "text": "风控研发组",
            "children": [{"id": "190", "type": "department", "text": "元方系统研发"}, {
              "id": "316",
              "type": "department",
              "text": "系统测试"
            }, {"id": "192", "type": "department", "text": "大数据平台"}, {
              "id": "356",
              "type": "department",
              "text": "元方系统研发（成都）"
            }, {"id": "380", "type": "department", "text": "智能平台（成都）"}]
          }, {
            "id": "328",
            "type": "department",
            "text": "决策科学组",
            "children": [{"id": "350", "type": "department", "text": "风控建模"}, {
              "id": "351",
              "type": "department",
              "text": "人工智能"
            }, {"id": "361", "type": "department", "text": "风控分析"}]
          }, {"id": "337", "type": "department", "text": "对外接口组"}]
        }, {
          "id": "174",
          "type": "department",
          "text": "O2-核心系统中心",
          "children": [{"id": "175", "type": "department", "text": "BI组"}, {
            "id": "176",
            "type": "department",
            "text": "财务系统组",
            "children": [{"id": "182", "type": "department", "text": "测试组"}, {
              "id": "314",
              "type": "department",
              "text": "开发组"
            }]
          }, {
            "id": "177",
            "type": "department",
            "text": "产品组",
            "children": [{"id": "183", "type": "department", "text": "资金合作组"}, {
              "id": "184",
              "type": "department",
              "text": "后端产品组"
            }]
          }, {"id": "381", "type": "department", "text": "信息平台组"}, {
            "id": "398",
            "type": "department",
            "text": "基础服务组",
            "children": [{"id": "10000026", "type": "department", "text": "test_add"}]
          }, {"id": "400", "type": "department", "text": "质量保障组"}, {
            "id": "407",
            "type": "department",
            "text": "支付与清结算开发组",
            "children": [{"id": "395", "type": "department", "text": "还款开发组"}, {
              "id": "394",
              "type": "department",
              "text": "放款开发组"
            }, {"id": "396", "type": "department", "text": "支付开发组"}, {
              "id": "397",
              "type": "department",
              "text": "清结算开发组"
            }]
          }]
        }, {
          "id": "385",
          "type": "department",
          "text": "O2-公共研发中心",
          "children": [{"id": "393", "type": "department", "text": "架构支持组"}, {
            "id": "392",
            "type": "department",
            "text": "测试组"
          }, {"id": "387", "type": "department", "text": "前端开发组"}, {
            "id": "240",
            "type": "department",
            "text": "IT支持组"
          }, {"id": "241", "type": "department", "text": "运维开发组"}, {
            "id": "220",
            "type": "department",
            "text": "值班组"
          }, {
            "id": "1033",
            "type": "department",
            "text": "创新产品部",
            "children": [{"id": "1045", "type": "department", "text": "前端组"}, {
              "id": "1044",
              "type": "department",
              "text": "创新项目组"
            }]
          }, {"id": "1034", "type": "department", "text": "质量保障部"}, {
            "id": "1035",
            "type": "department",
            "text": "基础架构部"
          }, {"id": "1032", "type": "department", "text": "信贷平台部"}, {
            "id": "1031",
            "type": "department",
            "text": "企业产品部"
          }, {"id": "1029", "type": "department", "text": "BI组"}]
        }, {
          "id": "234",
          "type": "department",
          "text": "O4-成都研发中心",
          "children": [{"id": "355", "type": "department", "text": "产品组"}, {
            "id": "237",
            "type": "department",
            "text": "核心系统研发组",
            "children": [{"id": "258", "type": "department", "text": "放款开发组"}, {
              "id": "261",
              "type": "department",
              "text": "基础服务组",
              "children": [{"id": "263", "type": "department", "text": "信息服务"}, {
                "id": "264",
                "type": "department",
                "text": "签约服务"
              }, {"id": "331", "type": "department", "text": "中间件3"}]
            }, {"id": "259", "type": "department", "text": "还款开发组"}, {
              "id": "260",
              "type": "department",
              "text": "支付开发组"
            }, {
              "id": "402",
              "type": "department",
              "text": "财务开发组",
              "children": [{"type": "employee", "text": "租户 1号", "id": "KN02241"}, {
                "type": "employee",
                "text": "租户 2号",
                "id": "KN02242"
              }]
            }]
          }, {
            "id": "332",
            "type": "department",
            "text": "风控系统研发组",
            "children": [{"id": "346", "type": "department", "text": "大数据开发组"}, {
              "id": "347",
              "type": "department",
              "text": "元方开发组"
            }, {"id": "377", "type": "department", "text": "模型开发组"}]
          }, {
            "id": "333",
            "type": "department",
            "text": "质量保障组",
            "children": [{"id": "334", "type": "department", "text": "核心系统质量组"}, {
              "id": "335",
              "type": "department",
              "text": "基础服务质量组"
            }, {"id": "336", "type": "department", "text": "线上质量组"}]
          }, {"id": "403", "type": "department", "text": "人力资源组"}, {
            "id": "404",
            "type": "department",
            "text": "行政组"
          }, {"id": "11009", "type": "department", "text": "oa2-test-001"}]
        }, {
          "id": "102",
          "type": "department",
          "text": "S1-财务中心",
          "children": [{
            "id": "103",
            "type": "department",
            "text": "会计核算组",
            "children": [{"id": "408", "type": "department", "text": "核算"}, {
              "id": "409",
              "type": "department",
              "text": "税务"
            }, {"id": "410", "type": "department", "text": "报表"}]
          }, {
            "id": "340",
            "type": "department",
            "text": "资金管理组",
            "children": [{"id": "341", "type": "department", "text": "出纳"}, {
              "id": "342",
              "type": "department",
              "text": "资金"
            }, {"id": "343", "type": "department", "text": "清算"}]
          }, {"id": "109", "type": "department", "text": "管理分析组"}, {"id": "118", "type": "department", "text": "内控内审组"}]
        }, {
          "id": "130",
          "type": "department",
          "text": "S2-人事行政中心",
          "children": [{
            "id": "131",
            "type": "department",
            "text": "人力资源组",
            "children": [{"id": "134", "type": "department", "text": "人事管理组"}, {
              "id": "136",
              "type": "department",
              "text": "HRBP"
            }]
          }, {"id": "132", "type": "department", "text": "行政组"}]
        }, {
          "id": "126",
          "type": "department",
          "text": "S3-政府合规中心",
          "children": [{"id": "127", "type": "department", "text": "公共事务组"}, {
            "id": "128",
            "type": "department",
            "text": "法务组"
          }, {"id": "256", "type": "department", "text": "资管业务组"}]
        }, {
          "id": "299",
          "type": "department",
          "text": "S4-机构业务中心",
          "children": [{"id": "300", "type": "department", "text": "机构合作组"}, {
            "id": "301",
            "type": "department",
            "text": "金融运营组"
          }]
        }, {
          "id": "140",
          "type": "department",
          "text": "S5-市场营销中心",
          "children": [{"id": "231", "type": "department", "text": "品牌公关组"}, {
            "id": "269",
            "type": "department",
            "text": "推广运营组"
          }]
        }, {
          "id": "386",
          "type": "department",
          "text": "S6-客服电销中心",
          "children": [{"id": "291", "type": "department", "text": "客服组"}, {
            "id": "358",
            "type": "department",
            "text": "电销组"
          }]
        }, {"id": "411", "type": "department", "text": "temp01"}, {
          "id": "1000002",
          "type": "department",
          "text": "芬睐研发中心-jc-modify",
          "children": [{
            "id": "1000003",
            "type": "department",
            "text": "芬睐研发中心-基础测试-2级子部门",
            "children": [{"id": "1000004", "type": "department", "text": "芬睐研发中心-基础测试-3级子部门"}]
          }]
        }, {"id": "1000022", "type": "department", "text": "芬睐研发中心-基础测试-Create2"}, {
          "id": "10000020",
          "type": "department",
          "text": "芬睐研发中心-jc-Job"
        }, {"id": "391", "type": "department", "text": "55"}, {
          "id": "283",
          "type": "department",
          "text": "腾梭",
          "children": [{
            "id": "363",
            "type": "department",
            "text": "研发部",
            "children": [{"id": "370", "type": "department", "text": "测试组"}, {
              "id": "371",
              "type": "department",
              "text": "风控产品组"
            }, {"id": "372", "type": "department", "text": "信贷系统组"}]
          }, {"id": "364", "type": "department", "text": "运营部"}, {
            "id": "365",
            "type": "department",
            "text": "人事行政部",
            "children": [{"id": "374", "type": "department", "text": "行政组"}, {
              "id": "375",
              "type": "department",
              "text": "人事组"
            }]
          }, {"id": "366", "type": "department", "text": "渠道管理部"}, {
            "id": "367",
            "type": "department",
            "text": "产品部"
          }, {"id": "368", "type": "department", "text": "售前部"}, {
            "id": "369",
            "type": "department",
            "text": "品牌公关部"
          }, {"id": "376", "type": "department", "text": "销售部"}, {"id": "382", "type": "department", "text": "财务法务中心"}]
        }]
      }], stepper: 1,
      abstract: '',
      dialog: false,
      title: '12313',
      image_url: 'https://t7.baidu.com/it/u=3893052629,2697609392&fm=193&f=GIF',
      content: null,
      templates:[],
    }
  },
  watch: {
    selection: function () {
      let receivers = [];
      this.selection.forEach(function (element) {
        receivers.push({id: element.id, type: element.type})
      });
      console.log(receivers)
    },
    channel: {
      handler() {
        if (this.channel != 0) {
          if (this.channel==1){
            this.templates= [
              {title: '模板1', description: '介绍文字1', content: '模板内容'},
              {
                title: '模板2',
                description: '介绍文字2',
                content: '<div class="mceTmpl"><span class="cdate">CDATE</span>，<span class="mdate">MDATE</span>，我的内容</div>'
              }
            ];
          }
          if (this.channel==2){
            this.templates= [
              {title: '模板3', description: '介绍文字1', content: '模板内容'},
              {
                title: '模板4',
                description: '介绍文字2',
                content: '<div class="mceTmpl"><span class="cdate">CDATE</span>，<span class="mdate">MDATE</span>，我的内容</div>'
              }
            ];
          }
          this.showEditor = false;
        }
      },
    },
  },
  created() {
    // this.showEditor=true;
  },
  methods: {
    editorInit(){
      console.log(this.content);
      this.showEditor=true;
      console.log(this.content);
    },
    preview(){
      this.content='<p><img src="https://oa-test-knzn02.oss-cn-hongkong.aliyuncs.com/graphics/2021-08-26/19603-0748bbaa51464434a46e7687a04ba7c3?OSSAccessKeyId=LTAI4Fv4DUBzjMg5qjWyudpN&amp;Expires=2260670239&amp;Signature=21gqqxqT4vYMhDhe%2F2IjD6nQbKA%3D" alt="" width="400" height="300" /></p>\n' +
          '<h2 style="text-align: center;">Welcome to the TinyMCE editor demo!</h2>\n';
      console.log(this.content)
    },
    notBeforeToday(date) {
      return date < moment().startOf('day').toDate();
    },
    store() {
      let receivers = [];
      this.selection.forEach(function (element) {
        receivers.push({id: element.id, type: element.type})
      });
      this.$axios.post('/graphics/message/store', {
        title: this.title,
        content: this.content,
        cover_image_url: this.image_url,
        receivers: receivers,
        send_time: this.send_time,
      }).then(response => {
        console.log(JSON.stringify(response.data));
        let resData = response.data;
        if (resData.code === 0) {
          this.$toast("保存成功！", {
            type: 'success',
            timeout: 2000
          });
          location.href = '/graphics/template/index'
        } else if (resData.code === 40 || resData.code === 1000004001) {
          this.$toast('保存出错：' + resData.message + '【' +
              resData.data.errors[Object.keys(resData.data.errors)[0]][0] + '】', {
            type: 'error',
          });
        } else {
          this.$toast('保存出错：' + resData.message, {
            type: 'error',
          });
        }
      })
          .catch(error => {
            console.log(error);
            this.$toast('保存出错：服务器出错！', {
              type: 'error',
              timeout: 2000,
            });
          })
          .finally(() => this.overlay = false);
    },
    publish() {

      this.showEditor=false;
      this.dialog = true;
    },
    getCoverImageUrl: function (url) {
      this.image_url = url
    },
    handleSearch: function (val) {
      if (val) {
        if (!this.allOpened) {
          this.lastOpen = this.open;
          this.allOpened = true;
          this.$refs.items.updateAll(true);
        }
      } else {
        this.$refs.items.updateAll(false);
        this.allOpened = false;
        this.open = this.lastOpen;
      }
    }
  },
}
</script>

<style>

.tox-tinymce {
  border-bottom:none !important;
  border-left:none !important;
  border-right:none !important;
}

</style>