<script lang="ts">
    import {Image, UserBlock} from "./style";

    import {currentSession} from "../../../../stores/auth";
    import DefaultUserImage from "../DefaultUserImage/DefaultUserImage.svelte";

</script>

<div class="{UserBlock} {Image}">
    {#if $currentSession.userData.image}
        <img class={Image} src={$currentSession.userData.image}
             alt={$currentSession.userData.login}/>
    {:else}
        <DefaultUserImage />
    {/if}
</div>