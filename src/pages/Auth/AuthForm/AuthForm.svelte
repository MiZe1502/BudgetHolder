<script lang="ts">
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
        <InputWithLabel label="Login" name="login" type="text"
                        bind:value={$currentRegData.login}/>
        <InputWithLabel label="Password" name="password" type="password"
                        bind:value={$currentRegData.password}/>
        <InputWithLabel label="Name" name="name" type="text"
                        bind:value={$currentRegData.name}/>
        <InputWithLabel label="Surname" name="surname" type="text"
                        bind:value={$currentRegData.surname}/>
    {:else}
        <InputWithLabel label="Login" name="login" type="text"
                        bind:value={$currentAuthData.login}/>
        <InputWithLabel label="Password" name="password" type="password"
                        bind:value={$currentAuthData.password}/>
    {/if}
    <div class="{ButtonsBlock}">
        {#if !isRegistrationForm}
        <Button title="Register" onClickHandler={() => onChangeForm(true)}
                secondary={true}
                        buttonClass={AuthButton}/>
        {:else}
        <Button title="Authorization" onClickHandler={() => onChangeForm(false)}
                secondary={true}
                        buttonClass={AuthButton}/>
        {/if}
        <Button title={isRegistrationForm ? "Save and login " : "Login"}
                onClickHandler={isRegistrationForm ? saveAndAuth : onAuth}
                buttonClass={AuthButton}/>
    </div>
</form>