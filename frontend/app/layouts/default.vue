<script></script>
<template>
  <div class="layout">
    <header>
      <div class="header-content">
        <h1>Todo List</h1>
      </div>
    </header>

    <main class="content">
      <slot />
    </main>

    <footer>
      <div class="footer-content">
        <h4>Todo List</h4>
        <p>Copyright@Alex Hsu</p>
        <p>alexhsu55150@gmail.com</p>
      </div>
    </footer>
  </div>
</template>
<style>
.layout{
    min-height: 100vh;
    display: flex;
    flex-direction: column;
    background-color: aliceblue;
}
header{
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 30px 0;
    .header-content{
      display: flex;
      flex-direction: column;
      justify-content: center;
    }
    .header-content h1{
       margin: 0;
       font-size: 50px;
       font-family: 'Franklin Gothic Medium', 'Arial Narrow', Arial, sans-serif;
       color: black; 
       text-align: center;
    }
}
.content{
    flex: 1;
    margin: 0;
    padding: 0;
}
footer{
    display: flex;
    flex-direction: column;
    justify-content: center;
    .footer-content{
        display: flex;
        flex-direction: column;
        align-items: center;
        padding-bottom: 20px;
    }
    .footer-content h4{
        margin: 2px;
    }
    .footer-content p{
        margin: 2px;
    }
}
</style>