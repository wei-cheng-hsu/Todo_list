<script setup>
import { ref, onMounted } from 'vue'
import {
  TrashIcon,
  PlusIcon
} from '@heroicons/vue/24/outline'

const config = useRuntimeConfig()
const loading = ref(false)
const todos = ref([])
const errorMessage = ref('')
const showAddTodo = ref(false)

const showEditTodo = ref(false)
const selectedTodo = ref(null)

const openTodo = (todo) => {
  selectedTodo.value = todo
  showEditTodo.value = true
}

const closeEditTodo = () => {
  showEditTodo.value = false
  selectedTodo.value = null
}

// 取得所有Todo資料(GetTodo)
 const fetchTodos = async () => {
   loading.value = true
   errorMessage.value = ''

   try {
     todos.value = await $fetch(
       `${config.public.apiBase}/api/todos`,
       {
        method: 'GET'
       }
     )
   } catch (error) {
     console.error('Failed to load todos:', error)
     errorMessage.value = 'Todo List 載入失敗'
   } finally {
     loading.value = false
   }
 }

// 新增Todo資料 (POST CreateTodo)
const addTodo = async (newTodo) => {
  errorMessage.value = ''

  try {
    await $fetch(
      `${config.public.apiBase}/api/todos`,
      {
        method: 'POST',

        body: {
          title: newTodo.title,
          description: newTodo.description,
          due_date: newTodo.dueDate
        }
      }
    )

    showAddTodo.value = false

    await fetchTodos()

  } catch (error) {
    console.error('Failed to create todo:', error)

    errorMessage.value = 'Todo 新增失敗'
  }
}

// 編輯Todo資料 (PUT EditTodo)
const editTodo = async(editedTodo) => {
  errorMessage.value = ''
  try{
    await $fetch(
      `${config.public.apiBase}/api/todos/${editedTodo.id}`,
      {
        method: 'PUT',
        body:{
          title: editedTodo.title,
          description: editedTodo.description,
          due_date: editedTodo.dueDate
        }
      }
    )

    showEditTodo.value = false
    selectedTodo.value = null

    await fetchTodos()
  }catch(error){
    console.error('Failed to update todo:', error)
    errorMessage.value = 'Todo修改失敗'
  }
}

// 完成Todo (PATCH FinishTodo)
const finishTodo = async (id) => {
  errorMessage.value = ''

  try{
    await $fetch(
      `${config.public.apiBase}/api/todos/${id}/finish`,
      {
        method: 'PATCH'
      }
    )

    showEditTodo.value = false
    selectedTodo.value = null

    await fetchTodos()

  }catch(error){
    console.error('Failed to finish todo:', error)
    errorMessage.value = 'Todo finish 更新失敗'
  }
}

// 刪除 Todo
const deleteTodo = async (id) => {
  errorMessage.value = ''

  try{
    await $fetch(
      `${config.public.apiBase}/api/todos/${id}`,
      {
        method: 'DELETE'
      }
    )

    await fetchTodos()
  }catch(error){
    console.error('Failed to delete todo:', error)
    errorMessage.value = 'Todo 刪除失敗'
  }
}

// 日期格式
const formatDate = (date) => {
  if (!date) {
    return ''
  }

  const dateObject = new Date(date)

  return [
    dateObject.getFullYear(),
    String(dateObject.getMonth() + 1).padStart(2, '0'),
    String(dateObject.getDate()).padStart(2, '0')
  ].join('.')
}

 onMounted(() => {
   fetchTodos()
 })
</script>

<template>
  <section class="todo-page">

    <div class="todo-container">

      <!-- 表格標題 -->
      <div class="todo-header">
        <div>Order</div>
        <div>Due Date</div>
        <div>Title</div>
        <div>Description</div>
        <div>Setup Date</div>
        <div>Status</div>
        <div></div>
      </div>


      <!-- 未來串 API 時才會真正使用 Loading -->
      <div
        v-if="loading"
        class="message"
      >
        Loading...
      </div>


      <!-- API Error -->
      <div
        v-else-if="errorMessage"
        class="message error"
      >
        {{ errorMessage }}
      </div>


      <!-- 沒有 Todo -->
      <div
        v-else-if="todos.length === 0"
        class="message"
      >
        目前沒有待辦事項
      </div>


      <!-- Todo List -->
      <div
        v-else
        class="todo-list"
      >

        <div
          v-for="(todo, index) in todos"
          :key="todo.id"
          class="todo-row"
          @click="openTodo(todo)"
        >

          <!-- Order -->
          <div class="order">
            {{ String(index + 1).padStart(2, '0') }}
          </div>

          <!-- Due Date -->
          <div>
            {{ formatDate(todo.due_date) }}
          </div>

          <!-- Title -->
          <div class="todo-title">
            {{ todo.title }}
          </div>

          <!-- Description -->
          <div class="todo-description">
            {{ todo.description }}
          </div>

          <!-- Setup Date -->
          <div>
            {{ formatDate(todo.setup_date) }}
          </div>

          <div class="todo-status">
            <span
              class="status"
              :class="`status-${todo.status}`"
            >
              {{ todo.status }}
            </span>
          </div>

          <!-- Delete -->
          <div class="delete-column">
            <button
                class="delete-button"
                type="button"
                aria-label="Delete todo"
                @click.stop="deleteTodo(todo.id)"
                >
                <TrashIcon class="delete-icon" />
            </button>
          </div>

        </div>

      </div>


      <!-- Add Button -->
      <button
        class="add-button"
        type="button"
        aria-label="Add todo"
        @click="showAddTodo = true"
        >
        <PlusIcon class="add-icon" />
      </button>

    </div>


    <!-- Add.vue -->
    <Add
      v-if="showAddTodo"
      @close="showAddTodo = false"
      @add="addTodo"
    />


    <!--Edit.vue-->
    <Edit
      v-if="showEditTodo && selectedTodo"
      :todo="selectedTodo"
      @close="closeEditTodo"
      @edit="editTodo"
      @finish="finishTodo"
    />

  </section>
</template>
<style scoped>
.todo-page {
  width: 100%;
  display: flex;
  justify-content: center;
  padding-left: 20px;
  padding-right: 20px;
  padding-top: 0;
  padding-bottom: 40px;
  box-sizing: border-box;
}

.todo-container {
  position: relative;

  width: 100%;
  max-width: 1300px;
  min-height: 650px;

  padding: 40px 24px 110px;

  border: 3px solid #222;
  border-radius: 40px;

  background-color: white;

  box-sizing: border-box;
}

/* Header */

.todo-header {
  display: grid;

  grid-template-columns:
    80px
    150px
    160px
    minmax(200px, 1fr)
    150px
    110px
    60px;

  align-items: center;
  gap: 15px;

  padding: 0 20px 22px;

  font-size: 22px;
  font-weight: 600;
}

/* Todo rows */

.todo-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.todo-row {
  display: grid;

  grid-template-columns:
    80px
    150px
    160px
    minmax(200px, 1fr)
    150px
    110px
    60px;

  align-items: center;

  gap: 15px;

  min-height: 70px;

  padding: 0 20px;

  border: 3px solid #222;
  border-radius: 18px;

  box-sizing: border-box;

  font-size: 19px;
}
.todo-title{
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}
.todo-description{
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}

.status-column {
  display: flex;
  justify-content: center;
  align-items: center;
}

.status {
  min-width: 70px;

  padding: 6px 10px;

  border-radius: 20px;

  text-align: center;

  font-size: 15px;
  font-weight: 600;

  text-transform: capitalize;
}
.status-undo {
  background-color: #e5e7eb;
  color: #374151;
}

.status-finish {
  background-color: #bbf7d0;
  color: #166534;
}

.status-delay {
  background-color: #fecaca;
  color: #991b1b;
}

/* Order */

.order {
  text-align: center;
}

/* Delete */

.delete-column {
  display: flex;
  justify-content: center;
}

.delete-button {
  display: flex;
  justify-content: center;
  align-items: center;

  padding: 6px;

  border: none;
  background: transparent;

  cursor: pointer;
}

.delete-icon {
  width: 26px;
  height: 26px;
}

.delete-button:hover {
  transform: scale(1.1);
}

/* Add */

.add-button {
  position: absolute;

  right: 45px;
  bottom: 35px;

  width: 72px;
  height: 72px;

  display: flex;
  justify-content: center;
  align-items: center;

  padding: 0;

  border: 3px solid #222;
  border-radius: 50%;

  background-color: white;

  cursor: pointer;

  transition:
    transform 0.2s ease,
    background-color 0.2s ease;
}

.add-icon {
  width: 45px;
  height: 45px;
}

.add-button:hover {
  transform: scale(1.08);
  background-color: #f5f5f5;
}

/* Status */

.message {
  padding: 50px;

  text-align: center;

  font-size: 18px;
}

.error {
  color: red;
}
</style>