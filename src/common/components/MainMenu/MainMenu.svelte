<script lang="typescript">
    import SingleMenuBlock from '../SingleMenuBlock/SingleMenuBlock.svelte'
    import menuConfig from './menuConfig'
    import { getCurrentPath } from '../../utils/utils'

    export let location
</script>


<style lang="scss">
    .MainMenu {
        position: fixed;
        width: 100%;
        box-shadow: 0 4px 2px -2px gray;
    }
</style>

<header class="MainMenu FlexHor PrimaryBackground">
    {#each menuConfig as menu}
        <SingleMenuBlock location={location} url={menu.url} title={menu.title} />
    {/each}
</header>