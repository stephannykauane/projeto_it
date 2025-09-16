<template>
  <div class="button-alterar-profile">
    <input class="profile-alterar" type="button" :value="text"/>
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

.button-alterar-profile{
    display: flex;
    justify-content: center;
    width: 100%;
    text-align: center;
    align-items: center;
}

.profile-alterar {
    display: flex;
    font-size: 0.95em;
    font-family: 'Konkhmer Sleokchher';
    color: #ffffff;
    background-color: #70885A;
    text-align: center;
    align-items: center;
    min-width: 150px;
}

.profile-alterar{
    border: 2px solid transparent;
    cursor: pointer;
    width: 100%;
    padding: 0.4em 2.5em;
    border-radius: 1em;
    transition: .3s;
}

.profile-alterar:hover{
    transition: .3s;
    background-color: #414e34d7;
}

@media screen and (max-width: 820px) {
    .profile-alterar{
        font-size: 0.8em;
        min-width: 140px;
    }
    
}
@media screen and (max-width: 769px) {
    .profile-alterar{
        font-size: 0.8em;
        min-width: 110px;
    }
    
}

@media screen and (max-width: 650px) {
    .profile-alterar{
    
        min-width: 100px;
    }
    
}

@media screen and (max-width: 600px) {
    .profile-alterar{
    
        min-width: 80px;
        padding: 0 1em;
    }
    
}

@media screen and (max-width: 580px) {
    .profile-alterar{
        min-width: 60px;
        padding: 0;
    }
    
}



@media screen and (max-width: 580px) {
    .profile-alterar{
        padding: 0.1em 0.5em;
        font-size: 0.6em;
        text-align: center;
        align-items: center;
    }
    
}
</style>


