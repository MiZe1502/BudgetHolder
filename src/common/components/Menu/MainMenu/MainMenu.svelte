<script lang="typescript">
    import { MainMenu, FlexHorCenter } from "./style";

    import SingleMenuBlock from '../SingleMenuBlock/SingleMenuBlock.svelte'
    
    import menuConfig from './menuConfig'
</script>

<header class="{MainMenu} {FlexHorCenter}">
    {#each menuConfig as menu}
        <SingleMenuBlock url={menu.url} title={menu.title} />
    {/each}
</header>