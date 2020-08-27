<script lang="ts">
    import {_} from 'svelte-i18n'
    import {style} from "./style";
    import {UserDataExtended} from "../../../../stores/auth";
    import {
        openedPopups,
        updatePopupInnerValidation
    } from "../../../../stores/popup";

    import InputWithLabel
        from "../../Inputs/InputWithLabel/InputWithLabel.svelte";
    import {onMount} from "svelte";


    const validateForm = (event: Event<HTMLInputElement>) => {
        //Dirty hack to change data before on:input
        if (event) {
            data[event.target.name] = event.target.value;
        }

        //validation
    }

    onMount(() => {
        validateForm(null);
    })

    $: formErrors = $openedPopups.filter((popup) => popup.uuid === outerPopupUuid).length > 0 && $openedPopups.filter((popup) => popup.uuid === outerPopupUuid)[0].innerValidationErrors;

    export let data: UserDataExtended
    export let outerPopupUuid: string = "";
</script>

<form class="{style.SideMinorPadding} {style.FlexVert} {style.Form}">
    <InputWithLabel label={$_("auth.labels.login")} name="login" type="text"
                    bind:value={data.login}/>
    <InputWithLabel label={$_("auth.labels.prev_password")} name="prev_password"
                    type="password"
                    bind:value={data.prevPassword}/>
    <InputWithLabel label={$_("auth.labels.password")} name="password"
                    type="password"
                    bind:value={data.password}/>
    <InputWithLabel label={$_("auth.labels.name")} name="name" type="text"
                    bind:value={data.name}/>
    <InputWithLabel label={$_("auth.labels.surname")} name="surname"
                    type="text"
                    bind:value={data.surname}/>
</form>