<script lang="typescript">
    import {_} from 'svelte-i18n'
    import {style} from "./style";

    import {currentSession} from "../../../../pages/Auth/auth";
    import SingleMenuBlock from '../SingleMenuBlock/SingleMenuBlock.svelte'

    import menuConfig from './menuConfig'
    import UserBlock from "../UserBlock/UserBlock.svelte";
</script>

<header class="{style.MainMenu} {style.FlexHorCenter}">
    <div class="{style.SideMainPadding} {style.FlexHorCenter} {style.Wrapper}">
        <div class="{style.FlexHorCenter}">
            {#each menuConfig as menu}
                <SingleMenuBlock url={menu.url} title={menu.title}/>
            {/each}
        </div>

        {#if $currentSession.userData}
            <UserBlock/>
        {/if}
    </div>
</header>