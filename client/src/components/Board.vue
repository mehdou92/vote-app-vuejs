<template>
  <div class="board">
    <a href="https://google.fr" @click.prevent="$alert($event.target.href)">Google</a>
    <h1 class="board-name">{{ name }}</h1>
    <div class="board-lists">
      <list :onNewCard="onNewCard" v-for="list in lists" v-bind="list" :key="list.id" />
    </div>
  </div>
</template>


<script>
// test
import List from './List.vue'
export default {
  name: 'Board',
  components: {
    list: List
  },
  data: () => ({
    draggedItem: null,
    lists: [
      {
        name: "List1",
        cards: [
          {
            name: "Card1"
          },
          {
            name: "Card2"
          },
          {
            name: "Card3"
          }
        ]
      },
      {
        name: "List2",
        cards: [
          {
            name: "Card1"
          },
          {
            name: "Card2"
          }          
        ]
      },
      {
        name: "List3",
        cards: []
      }
    ],
    name: "Board1"
  }),
  provide: function() {
    return {
      onDrag: this.onDrag,
      onDrop: this.onDrop,
    };
  },

  methods: {
    onNewCard: function(card, into) {
      this.lists = this.lists.map(list => {
        if(list.name != into.name) return list;
        return { 
          ...list, 
          cards: [card, ...list.cards]
        }
      });
    },
    onDrag: function(card, list){
      this.draggedItem = {
        card,
        list
      };
    },
    onDrop: function(into){
      this.list = this.lists.map(list => {
        if(list.name === into){
          list.cards = [this.draggedItem.card, ...list.cards];
        }
        if(this.draggedItem.list.name === list.name){
          list.cards = list.cards.filter(
            card => card.name !== this.draggedItem.card.name
          );
        }
      });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  div.board {
    display: flex;
    flex-direction: column;
    flex: 1;
    background-color: red; 
  }
  div.board-name {
    background-color: green;
  }
  div.board-lists {
    flex: 1;
    display: flex;
    flex-direction: row;
    background-color: yellow;
    width: 100%;
    overflow-x: scroll;
  }
</style>