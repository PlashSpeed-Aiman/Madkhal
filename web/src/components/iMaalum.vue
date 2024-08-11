<script setup lang="ts">

import {ref} from "vue";

//get value of current year in javascript

const sessionYear = ref(`${new Date().getFullYear()}/${new Date().getFullYear()+1}`)
const semester = ref('1')
function GetFinance() {
  fetch('http://localhost:8080/finance', {
    method: 'GET',
    headers: {
      'Access-Control-Allow-Origin': '*',
      'Content-Type': 'text/plain',
    },

  }).then(res => res.json()).then(data => {
    alert(data.Message)
  })
}
function GetResult() {
  fetch('http://localhost:8080/result', {
    method: 'POST',
    headers: {
      'Access-Control-Allow-Origin': '*',
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      year: sessionYear.value,
      semester: semester.value
    })

  }).then(res => res.json()).then(data => {
    alert(data.Message)
  })
}
function GetExamSlip() {
  fetch('http://localhost:8080/es', {
    method: 'GET',
    headers: {
      'Access-Control-Allow-Origin': '*',
      'Content-Type': 'text/plain',
    },

  }).then(res => res.json()).then(data => {
    alert(data.Message)
  })
}
function GetConfirmationSlip() {
  fetch('/cs', {
    method: 'POST',
    headers: {
      'Access-Control-Allow-Origin': '*',
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      year: sessionYear.value,
      semester: semester.value
    })

  }).then(res => res.json()).then(data => {
    alert(data.Message)
  })
}
</script>

<template>
  <div class=" mx-auto my-auto border p-6 rounded-lg bg-slate-100">
    <div class="flex flex-col items-center">

    <h1 class="text-3xl my-2">i-Maalum</h1>
    <p >Make sure to setup your credentials in the settings page</p>
    <div class="my-2">
      <form>
        <p class="font-semibold">Session (eg 2021/2022)</p>
        <input v-model="sessionYear" class="border-2 rounded px-1 py-1"/>
        <p class="font-semibold">Semester</p>
        <select v-bind="semester" class="border-2 rounded px-1 py-1 w-full">
          <option>1</option>
          <option>2</option>
          <option>3</option>
        </select>
      </form>
      <div class="flex flex-col justify-center gap-1.5 my-2">
        <button v-on:click="GetFinance" class="border border-[#455A64] bg-[#CFD8DC] text-[#212121] transition-colors duration-100 rounded p-2 hover:bg-[#607D8B] hover:text-white text-center">Finance</button>
        <button v-on:click="GetResult" class="border border-[#455A64] bg-[#CFD8DC] text-[#212121] transition-colors duration-100 rounded p-2 hover:bg-[#607D8B] hover:text-white text-center">Result</button>
        <button v-on:click="GetConfirmationSlip" class="border border-[#455A64] bg-[#CFD8DC] text-[#212121] transition-colors duration-100 rounded p-2 hover:bg-[#607D8B] hover:text-white text-center">Course Confirmation Slip</button>
        <button v-on:click="GetExamSlip" class="border border-[#455A64] bg-[#CFD8DC] text-[#212121] transition-colors duration-100 rounded p-2 hover:bg-[#607D8B] hover:text-white text-center">Exam Slip</button>
      </div>
    </div>
    </div>

  </div>
</template>

<style scoped>

</style>