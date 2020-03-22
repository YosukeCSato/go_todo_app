<template>
  <div>
    <label v-for="label in options" :key="label">
      <input type="radio" v-model="current" v-bind:value="label.value" />{{
        label.label
      }}
    </label>
    <p>
      品名:<input
        type="text"
        name="productName"
        v-model="productName"
        placeholder="品名を入力してください"
      />
    </p>
    <p>
      メモ:<input
        type="text"
        name="productMemo"
        v-model="productMemo"
        placeholder="品名を入力してください"
      />
    </p>

    <button v-on:click="doAddProduct" v-bind:disabled="!isEntered">追加</button>
    <hr />
    <table>
      <thead v-pre>
        <tr>
          <th>No</th>
          <th>商品名</th>
          <th>メモ</th>
          <th>状態</th>
          <th>削除</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in computedProducts" :key="item">
          <td>{{ index + 1 }}</td>
          <td>{{ item.name }}</td>
          <td>{{ item.memo }}</td>
          <td>
            <button v-on:click="doChangeProductState(item)">
              {{ labels[item.state] }}
            </button>
          </td>
          <td>
            <button v-on:click="doDeleteProduct(item)">
              削除
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import axios from 'axios';

export default {
  data() {
    return {
      products: [],
      productName: '',
      productMemo: '',
      current: -1,
      options: [
        { value: -1, label: 'すべて' },
        { value: 0, label: '未購入' },
        { value: 1, label: '購入済み' },
      ],
      isEntered: false,
    };
  },

  computed: {
    labels() {
      return this.options.reduce(function(a, b) {
        return Object.assign(a, { [b.value]: b.label });
      }, {});
    },

    computedProducts() {
      return this.products.filter(function(el) {
        var option = this.current < 0 ? true : this.current === el.state;
        return option;
      }, this);
    },

    validate() {
      var isEnteredProductName = 0 < this.productName.length;
      this.isEntered = isEnteredProductName;
      return isEnteredProductName;
    },
  },

  created() {
    this.doFetchAllProducts();
  },

  methods: {
    doFetchAllProducts() {
      axios.get('http://localhost:8080/fetchAllProducts').then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー');
        } else {
          var resultProducts = response.data;
          this.products = resultProducts;
        }
      });
    },
  },
};
</script>
