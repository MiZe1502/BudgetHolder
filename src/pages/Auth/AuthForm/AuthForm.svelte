<script lang="ts">
    import {_} from 'svelte-i18n'
    import {FlexVert, Form, ButtonsBlock, AuthButton} from "./style";

    import InputWithLabel
        from "../../../common/components/Inputs/InputWithLabel/InputWithLabel.svelte";
    import {
        currentAuthData,
        currentRegData,
        mockAuthorize, mockSaveAndAuthorize
    } from "../../../stores/auth";
    import Button
        from "../../../common/components/Buttons/Button/Button.svelte";
    import {navigate} from "svelte-routing";
    import routes from "../../../common/utils/routes"

    let isRegistrationForm = false;

    const onAuth = (event: Event<HTMLFormElement>) => {
        event.preventDefault();
        const error = mockAuthorize($currentAuthData);

        if (!error) {
            navigate(routes.budget, {replace: true});
        }
    }

    const saveAndAuth = (event: Event<HTMLFormElement>) => {
        event.preventDefault();
        const error = mockSaveAndAuthorize($currentRegData);

        if (!error) {
            navigate(routes.budget, {replace: true});
        }
    }

    const onChangeForm = (isRegistration: boolean) => {
        isRegistrationForm = isRegistration;
    }

</script>

<form class="{Form} {FlexVert}">
    {#if isRegistrationForm}
        <InputWithLabel label={$_("auth.labels.login")} name="login" type="text"
                        bind:value={$currentRegData.login}/>
        <InputWithLabel label={$_("auth.labels.password")} name="password"
                        type="password"
                        bind:value={$currentRegData.password}/>
        <InputWithLabel label={$_("auth.labels.name")} name="name" type="text"
                        bind:value={$currentRegData.name}/>
        <InputWithLabel label={$_("auth.labels.surname")} name="surname"
                        type="text"
                        bind:value={$currentRegData.surname}/>
    {:else}
        <InputWithLabel label={$_("auth.labels.login")} name="login" type="text"
                        bind:value={$currentAuthData.login}/>
        <InputWithLabel label={$_("auth.labels.password")} name="password"
                        type="password"
                        bind:value={$currentAuthData.password}/>
    {/if}
    <div class="{ButtonsBlock}">
        {#if !isRegistrationForm}
        <Button title={$_("auth.buttons.register")}
                onClickHandler={() => onChangeForm(true)}
                secondary={true}
                        buttonClass={AuthButton}/>
        {:else}
        <Button title={$_("auth.buttons.auth")}
                onClickHandler={() => onChangeForm(false)}
                secondary={true}
                        buttonClass={AuthButton}/>
        {/if}
        <Button title={isRegistrationForm ? $_("auth.buttons.save_and_login") : $_("auth.labels.login")}
                onClickHandler={isRegistrationForm ? saveAndAuth : onAuth}
                buttonClass={AuthButton}/>
    </div>
</form>