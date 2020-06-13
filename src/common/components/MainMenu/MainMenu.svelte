<script lang="typescript">
    import SingleMenuBlock from '../SingleMenuBlock/SingleMenuBlock.svelte'
    import menuConfig from './menuConfig'
</script>


<style>
    .MainMenu {
        /* position: fixed; */
        width: 100%;
        box-shadow: 0 4px 2px -2px gray;
        margin-bottom: 40px;
    }
</style>

<header class="MainMenu FlexHor PrimaryBackground">
    {#each menuConfig as menu}
        <SingleMenuBlock url={menu.url} title={menu.title} />
    {/each}
</header>