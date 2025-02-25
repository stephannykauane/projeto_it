<template>
        <nav class="container-menu">
            <ul class="menu">
                <li><a href="#">Perfil</a></li>
                <li><a href="#">Calculadora</a></li>
                <li><a href="#">Painel Profissional</a></li> 
            </ul>
        </nav>
</template>


<style scoped>

.container-menu{
    width: 100%;
    display: flex;
    justify-content:center;
    border: 2px solid transparent;
    border-radius: 10em;

}

.menu{
    list-style: none;
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;

}

nav{
    background-color:#70885A;    
}
.menu a{
    display: flex;
    justify-content: center;  
    width: 100%;
    padding: 1em 8em;
    border: 2px solid transparent;
    border-radius: 10em;
    text-decoration: none;
    font-family: 'Konkhmer Sleokchher';
    text-transform: uppercase;
    color: rgba(202, 254, 157, 0.8);
    transition: .3s;
}

.menu a:hover{
    color: rgb(202, 254, 157);
    transition: .3s;
}
</style>