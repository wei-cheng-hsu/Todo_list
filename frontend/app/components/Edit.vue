<script setup>
import { reactive, computed } from 'vue';
import {
    XMarkIcon,
    CalendarDaysIcon
} from '@heroicons/vue/24/outline'

const props = defineProps({
    todo: {
        type: Object,
        required: true
    }
})

const emit = defineEmits([
    'close',
    'edit',
    'finish'
])

const form = reactive({
  title: props.todo.title,
  description: props.todo.description,
  dueDate: props.todo.due_date ? props.todo.due_date.slice(0, 10) : ''
})

const titleLength = computed(() => {
  return form.title.length
})

const descriptionLength = computed(() => {
  return form.description.length
})

const editTodo = () => {
  if (!form.title.trim()) {
    alert('Please enter title.')
    return
  }

  if (!form.dueDate) {
    alert('Please select due date.')
    return
  }

  emit('edit', {
    id: props.todo.id,
    title: form.title.trim(),
    description: form.description.trim(),
    dueDate: form.dueDate
  })
}

const finishTodo = () => {
  emit('finish', props.todo.id)
}

const closeModal = () => {
  emit('close')
}
</script>

<template>
  <div
    class="modal-overlay"
    @click.self="closeModal"
  >
    <div class="edit-modal">

      <!-- Close -->
      <button
        class="close-button"
        type="button"
        aria-label="Close"
        @click="closeModal"
      >
        <XMarkIcon class="close-icon" />
      </button>

      <form
        class="todo-form"
        @submit.prevent
      >

        <!-- Title -->
        <div class="form-group">
          <label for="edit-title">
            Title:
          </label>

          <div class="input-wrapper">
            <input
              id="edit-title"
              v-model="form.title"
              type="text"
              maxlength="50"
            >

            <span class="character-count">
              {{ titleLength }}/50
            </span>
          </div>
        </div>

        <!-- Description -->
        <div class="form-group">
          <label for="edit-description">
            Description:
          </label>

          <div class="textarea-wrapper">
            <textarea
              id="edit-description"
              v-model="form.description"
              maxlength="1000"
            />

            <span class="character-count description-count">
              {{ descriptionLength }}/1000
            </span>
          </div>
        </div>

        <!-- Due Date -->
        <div class="form-group due-date-group">
          <label for="edit-due-date">
            Due Date:
          </label>

          <div class="date-wrapper">
            <input
              id="edit-due-date"
              v-model="form.dueDate"
              type="date"
            >

            <CalendarDaysIcon class="calendar-icon" />
          </div>
        </div>

        <!-- Buttons -->
        <div class="button-group">

          <button
            class="finish-button"
            type="button"
            @click="finishTodo"
          >
            Finish
          </button>

          <button
            class="edit-button"
            type="button"
            @click="editTodo"
          >
            Edit
          </button>

          <button
            class="cancel-button"
            type="button"
            @click="closeModal"
          >
            Cancel
          </button>

        </div>

      </form>

    </div>
  </div>
</template>

<style scoped>
.modal-overlay {
  position: fixed;
  inset: 0;

  display: flex;
  justify-content: center;
  align-items: flex-start;

  padding: 30px;

  background-color: rgba(0, 0, 0, 0.35);

  z-index: 1000;

  box-sizing: border-box;
  overflow-y: auto;
}

.edit-modal {
  position: relative;

  width: 100%;
  max-width: 850px;

  min-height: 760px;

  padding: 75px 60px 45px;

  border: 2px solid #222;
  border-radius: 28px;

  background-color: white;

  box-sizing: border-box;

  margin: auto 0;
}

.close-button {
  position: absolute;

  top: 22px;
  right: 28px;

  width: 40px;
  height: 40px;

  display: flex;
  justify-content: center;
  align-items: center;

  padding: 0;

  border: none;
  background: transparent;

  cursor: pointer;
}

.close-icon {
  width: 32px;
  height: 32px;
}

.todo-form {
  height: 100%;

  display: flex;
  flex-direction: column;
}

.form-group {
  display: flex;
  flex-direction: column;

  gap: 12px;

  margin-bottom: 45px;
}

.form-group label {
  font-size: 19px;
  font-weight: 500;
}

.input-wrapper {
  position: relative;
}

.input-wrapper input {
  width: 100%;
  height: 50px;

  padding: 0 75px 0 16px;

  border: 2px solid #222;
  border-radius: 12px;

  font-size: 17px;

  box-sizing: border-box;
}

.textarea-wrapper {
  position: relative;
}

.textarea-wrapper textarea {
  width: 100%;
  height: 350px;

  padding: 18px 20px 45px;

  border: 2px solid #222;
  border-radius: 26px;

  font-family: inherit;
  font-size: 17px;

  resize: none;

  box-sizing: border-box;
}

.character-count {
  position: absolute;

  right: 16px;
  bottom: 13px;

  color: #999;

  font-size: 16px;

  pointer-events: none;
}

.description-count {
  right: 20px;
  bottom: 18px;
}

.due-date-group {
  width: 55%;
}

.date-wrapper {
  position: relative;
}

.date-wrapper input {
  width: 100%;
  height: 48px;

  padding: 0 45px 0 15px;

  border: 2px solid #222;
  border-radius: 10px;

  font-size: 17px;

  box-sizing: border-box;
}

.calendar-icon {
  position: absolute;

  top: 50%;
  right: 12px;

  width: 22px;
  height: 22px;

  transform: translateY(-50%);

  pointer-events: none;
}

.date-wrapper input[type="date"]::-webkit-calendar-picker-indicator {
  position: absolute;

  right: 10px;

  width: 25px;
  height: 25px;

  opacity: 0;

  cursor: pointer;
}

.button-group {
  margin-top: auto;

  display: flex;
  justify-content: flex-end;

  gap: 20px;
}

.finish-button,
.edit-button,
.cancel-button {
  min-width: 120px;

  padding: 11px 25px;

  border: 2px solid #222;
  border-radius: 12px;

  font-size: 17px;

  cursor: pointer;

  transition:
    transform 0.2s ease,
    opacity 0.2s ease;
}

.finish-button {
  background-color: #bbf7d0;
}

.edit-button {
  background-color: #bfdbfe;
}

.cancel-button {
  background-color: #ffbaba;
}

.finish-button:hover,
.edit-button:hover,
.cancel-button:hover {
  transform: translateY(-2px);
  opacity: 0.85;
}

@media (max-width: 700px) {
  .modal-overlay {
    padding: 15px;
  }

  .edit-modal {
    min-height: auto;

    padding: 65px 25px 30px;
  }

  .textarea-wrapper textarea {
    height: 250px;
  }

  .due-date-group {
    width: 100%;
  }

  .button-group {
    margin-top: 20px;
    flex-direction: column;
  }

  .finish-button,
  .edit-button,
  .cancel-button {
    width: 100%;
  }
}
</style>