<script lang="ts">
    import SingleMenuBlock from "../SingleMenuBlock/SingleMenuBlock.svelte";
    //import colors from "../../utils/colors"
    import menuConfig from "./menuConfig";
</script>


<style lang="scss">
    .MainMenu {
        position: fixed;
        width: 100%;
        border:  1px solid black;
    }

    .HorFlex {
        display: flex;
        flex-direction: row;
        align-items: center;
    }
</style>

<header class="MainMenu HorFlex">
    {#each menuConfig as menu}
        <SingleMenuBlock url={menu.url} title={menu.title} />
    {/each}
</header>