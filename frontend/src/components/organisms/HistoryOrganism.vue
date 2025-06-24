<script setup lang="ts">


import HeaderAndNavMolecule from '../molecules/HeaderAndNavMolecule.vue';
import CheckBoxMolecule from '../molecules/CheckBoxMolecule.vue';

</script>

<template>
    <div class="container-history-organism">
        <div class="header-and-nav">
            <HeaderAndNavMolecule />
        </div>
        <div class="history-things">
            <div class="history-text">
                <div class="history-text-icon">
                    <div>
                        <h1>HISTÓRICO</h1>
                    </div>
                    <div class="history-icon">
                        <img class="icon-history" src="../../../src/assets/history-svgrepo-com.svg">
                    </div>
                </div>
            </div>
        </div>
        <div class="checkbox-size">
            <CheckBoxMolecule />
        </div>
    </div>
</template>


<style scoped>
.container-history-organism {
    display: flex;
    width: 100%;
    justify-content: center;
    flex-direction: column;
}

.history-text {
    font-family: 'Konkhmer Sleokchher';
}

.checkbox-size {
    display: flex;
    justify-content: center;
}

.history-text h1 {
    margin-top: 0.1em;
    font-size: 1.8em;

}

.filter {
    display: flex;
    justify-content: center;
    margin: 1em 4em;
    position: relative;
    z-index: 5;
}


.history-text-icon {
    width: 100%;
    display: flex;
    justify-content: flex-start;
    flex-direction: row;
    align-items: center;
}

.icon-history {
    width: 3.5%;
    margin-top: 0.2em;
}

.icon-filter {
    width: 2.5%;
}


.history-things {
    display: flex;
    justify-content: flex-start;
    flex-direction: column;
    padding-left: 16.3em;
    margin-bottom: 1em;
    color: #ffffff;
}

.history-icon {
    display: flex;
    justify-content: flex-start;
    margin-left: 0.5em;
    margin-bottom: 0.3em;
}

.icon-2 {
    display: flex;
    justify-content: flex-start;
    margin-left: 0.5em;
}

.filter-text-icon {
    width: 100%;
    text-align: left;
    display: flex;
    justify-content: flex-start;
    flex-direction: row;
    align-items: center;
    font-size: 1.5em;
    margin-top: 1em;
}

.select {
    width: 23.4%;
    margin-top: 3em;
    margin-left: -0.5em;
}

::v-deep(.profile-alterar) {
    background-color: #4f613f;
    padding: 0.4em 1em;
    margin-left: 17.8em;

}

::v-deep(.profile-alterar:hover) {
    background-color: #344129;

}


@media screen and (max-width:1280px) {
    .history-things {
        display: flex;
        justify-content: flex-start;
        flex-direction: column;
        padding-left: 8em;
        margin-bottom: 1em;
        margin-top:2em;
        color: #ffffff;
    }

    


}

@media screen and (max-width: 769px) {
    .history-things {
        display: flex;
        justify-content: flex-start;
        flex-direction: column;
        padding-left: 5.5em;
        margin-bottom: 1em;
        margin-top:2em;
        color: #ffffff;
    }

    .history-text {
        font-size: 0.7em;
    }
    
    .icon-history {
        width: 3.4%;
        margin-top: 0.3em;
    }


}

@media screen and (max-width: 540px) {

    .history-things {
        display: flex;
        justify-content: flex-start;
        flex-direction: column;
        padding-left: 3.5em;
        margin-bottom: 1em;
        margin-top:2em;
        color: #ffffff;
    }


    .history-text {
        font-size: 0.7em;
    }

    .icon-history {
        width: 6.5%;
        margin-top: 0.3em;
    }


}


@media screen and (max-width: 480px) {
    .history-things {
        display: flex;
        justify-content: flex-start;
        flex-direction: column;
        padding-left: 1em;
        margin-bottom: 1em;
        margin-top:2em;
        color: #ffffff;
    }

    .history-text {
        font-size: 0.7em;
    }

    .icon-history {
        width: 6.5%;
        margin-top: 0.3em;
    }


}
</style>