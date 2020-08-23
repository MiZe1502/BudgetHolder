<script lang="ts">
    import {
        Image,
        UserBlock,
        FlexHorCenter,
        Popup,
        PopupItem,
        Font312Black
    } from "./style";
    import ClickOutside from "svelte-click-outside";

    import {currentSession} from "../../../../stores/auth";
    import DefaultUserImage from "../DefaultUserImage/DefaultUserImage.svelte";


    let isUserPopupOpened = false;

    const onUserImageClick = () => {
        isUserPopupOpened = !isUserPopupOpened;
    }

    const onClickOutside = () => {
        isUserPopupOpened = false;
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
            <div class="{Popup}">
                <div class="{PopupItem} {FlexHorCenter} {Font312Black}">
                    Settings
                </div>
                <div class="{PopupItem} {FlexHorCenter} {Font312Black}">
                    Logout
                </div>
            </div>
        {/if}
    </ClickOutside>
</div>

