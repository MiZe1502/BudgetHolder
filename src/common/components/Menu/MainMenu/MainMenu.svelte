<script lang="typescript">
    import {_} from 'svelte-i18n'
    import {style} from "./style";

    import {authStatus} from "../../../../stores/auth";
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

        {#if $authStatus}
            <UserBlock/>
        {/if}
    </div>
</header>