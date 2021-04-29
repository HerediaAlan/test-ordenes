<template>
    <div id="homeContainer">
        <Menu />
        <router-view></router-view>
    </div>
</template>

<script>
import Menu from "@/components/Menu";

export default {
    name: "Home",
    component: Menu
};
</script>

<style>
#mainBody {
    position: absolute;
    top: 50%;
    left: 50%;
}
</style>