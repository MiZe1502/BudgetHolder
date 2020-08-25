<script lang="typescript">
    import {_} from 'svelte-i18n'
    import {
        MainMenu,
        FlexHorCenter,
        SideMainPadding,
        Wrapper
    } from "./style";

    import {authStatus} from "../../../../stores/auth";
    import SingleMenuBlock from '../SingleMenuBlock/SingleMenuBlock.svelte'

    import menuConfig from './menuConfig'
    import UserBlock from "../UserBlock/UserBlock.svelte";
</script>

<header class="{MainMenu} {FlexHorCenter}">
    <div class="{SideMainPadding} {FlexHorCenter} {Wrapper}">
        <div class="{FlexHorCenter}">
            {#each menuConfig as menu}
                <SingleMenuBlock url={menu.url} title={menu.title}/>
            {/each}
        </div>

        {#if $authStatus}
            <UserBlock/>
        {/if}
    </div>
</header>