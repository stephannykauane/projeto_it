<template>
    <div class="button-exportar-excel">
      <input class="exportar-excel" type="button" :value="text"/>
      <img class="excel-icon" src="../../../src/assets/excel.svg">
    </div>
  </template>
  
  <script>
  export default {
      props: {
          text: String,
      }
  };
  </script>
  
  <style>
  
  .button-exportar-excel{
      display: flex;
      justify-content: center;
      width: 100%;
      padding: 0.7em;
      text-align: center;
      flex-direction: row;
      align-items: center;
      background-color: #70885A;
      border-radius: 0.8em;
      transition: .3s;
  }

  .excel-icon{
    width: 30px;
  }
  
  .exportar-excel{
      display: flex;
      font-size: 0.9em;
      font-family: 'Konkhmer Sleokchher';
      color: #ffffff;
      text-align: center;
      background-color: transparent;
      align-items: center;
      min-width: 150px;
 
  }
  
  .exportar-excel{
      border: 2px solid transparent;
      cursor: pointer;
      width: 50%;
  }
  
  .button-exportar-excel:hover{
      background-color: #414e34;
      transition: .3s;
  }

  

  @media screen and (max-width: 820px) {
      .exportar-excel{
          font-size: 0.8em;
          min-width: 140px;
      }

      .excel-icon {
        width: 20px;
      }

      .button-exportar-excel {
        padding: 0.64em;
      }
      
  }
  
  
  
  @media screen and (max-width: 769px) {
      .exportar-excel{
          font-size: 0.8em;
          min-width: 103px;
      }

      .excel-icon {
        width: 20px;
      }

      .button-exportar-excel {
        padding: 0.64em;
      }
      
  }
  
  
  </style>
  
  
  