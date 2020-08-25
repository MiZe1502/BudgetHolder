<script lang="ts">
    import {_} from 'svelte-i18n'

    import {
        Image,
        UserBlock,
        FlexHorCenter,
        PopupItem,
        Font312Black
    } from "./style";
    import ClickOutside from "svelte-click-outside";
    import routes from "../../../../common/utils/routes";

    import {
        authStatusKey,
        sessionKey,
        currentSession, setCurrentSession, setAuthStatus
    } from "../../../../stores/auth";
    import DefaultUserImage from "../DefaultUserImage/DefaultUserImage.svelte";
    import {removeDataFromLocalStorageByKey} from "../../../utils/localStorage";
    import {navigate} from "svelte-routing";
    import Dropdown from "../../ElementsAndBlocks/Dropdown/Dropdown.svelte";


    let isUserPopupOpened = false;

    const onUserImageClick = () => {
        isUserPopupOpened = !isUserPopupOpened;
    }

    const onClickOutside = () => {
        isUserPopupOpened = false;
    }

    const onLogout = () => {
        setAuthStatus(false);
        setCurrentSession({});

        removeDataFromLocalStorageByKey(sessionKey);
        removeDataFromLocalStorageByKey(authStatusKey);

        navigate(routes.auth, {replace: true})
    }

</script>

<div on:click={onUserImageClick} class="{UserBlock} {Image} {FlexHorCenter}"
     title={$currentSession.userData.login}>
    <ClickOutside on:clickoutside={onClickOutside}>

        {#if $currentSession.userData.image}
            <img class={Image}
                 src={$currentSession.userData.image}
                 alt={$currentSession.userData.login}/>
        {:else}
            <DefaultUserImage/>
        {/if}
        {#if isUserPopupOpened}
            <Dropdown width={200}>
                <div class="{PopupItem} {FlexHorCenter} {Font312Black}">
                    {$_("user.menu.settings")}
                </div>
                <div on:click={onLogout}
                     class="{PopupItem} {FlexHorCenter} {Font312Black}">
                    {$_("user.menu.logout")}
                </div>
            </Dropdown>
        {/if}
    </ClickOutside>
</div>

