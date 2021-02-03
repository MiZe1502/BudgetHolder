<script lang="ts">
    import {_} from 'svelte-i18n'

    import {style} from "./style";
    import ClickOutside from "svelte-click-outside";
    import routes from "../../../../common/utils/routes";

    import {
        authStatusKey,
        sessionKey,
        currentSession, setCurrentSession, setAuthStatus, logout
    } from "../../../../pages/Auth/auth";
    import DefaultUserImage from "../DefaultUserImage/DefaultUserImage.svelte";
    import {removeDataFromLocalStorageByKey} from "../../../utils/localStorage";
    import {navigate} from "svelte-routing";
    import Dropdown from "../../ElementsAndBlocks/Dropdown/Dropdown.svelte";
    import UserSettingsPopup
        from "../UserSettingsPopup/UserSettingsPopup.svelte";


    let isUserPopupOpened = false;
    let isUserSettingsPopupOpened = false;

    const onUserImageClick = () => {
        isUserPopupOpened = !isUserPopupOpened;
    }

    const onClickOutside = () => {
        isUserPopupOpened = false;
    }

    const onUserSettingsClick = () => {
        isUserSettingsPopupOpened = true;
    }

    const onLogout = async () => {
        await logout()
        navigate(routes.auth, {replace: true})
    }

</script>

<div on:click={onUserImageClick}
     class="{style.UserBlock} {style.Image} {style.FlexHorCenter}"
     title={$currentSession.userData.login}>
    <ClickOutside on:clickoutside={onClickOutside}>

        {#if $currentSession.userData.image}
            <img class={style.Image}
                 src={$currentSession.userData.image}
                 alt={$currentSession.userData.login}/>
        {:else}
            <DefaultUserImage/>
        {/if}
        {#if isUserPopupOpened}
            <Dropdown width={200}>
                <div on:click={onUserSettingsClick}
                     class="{style.PopupItem} {style.FlexHorCenter} {style.Font312Black}">
                    {$_("user.menu.settings")}
                </div>
                <div on:click={onLogout}
                     class="{style.PopupItem} {style.FlexHorCenter} {style.Font312Black}">
                    {$_("user.menu.logout")}
                </div>
            </Dropdown>
        {/if}
    </ClickOutside>
</div>
{#if isUserSettingsPopupOpened}
    <UserSettingsPopup data={$currentSession.userData} bind:isPopupOpened={isUserSettingsPopupOpened}/>
{/if}

