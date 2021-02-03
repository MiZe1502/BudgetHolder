<script lang="ts">
    import {_} from 'svelte-i18n'
    import {
        style
    } from "./style";
    import {
        clearValidationResults,
        validationResults
    } from "../../Budget/purchases";

    import InputWithLabel
        from "../../../common/components/Inputs/InputWithLabel/InputWithLabel.svelte";
    import {
        authorize,
        clearAuthAndRegData,
        currentSession,
        currentAuthData,
        currentRegData, getUserData, registerUser
    } from "../auth";
    import Button
        from "../../../common/components/Buttons/Button/Button.svelte";
    import {navigate} from "svelte-routing";
    import routes from "../../../common/utils/routes"
    import {validationRules} from "./validationRules";
    import {updateValidationResults} from "../../Budget/purchases";

    let isRegistrationForm = false;

    const validateForm = (currentData) => {
        for (let rule of validationRules) {
            const isInvalid = rule.validator(currentData);
            if (isInvalid) {
                updateValidationResults(rule.message)
            }
        }
    }

    const onAuth = async (event: Event<HTMLFormElement>) => {
        event.preventDefault();

        clearValidationResults();

        validateForm($currentAuthData);

        if ($validationResults.length > 0) {
            return;
        }

        const error = await authorize($currentAuthData);

        await getUserData($currentAuthData.login)

        updateValidationResults(error);

        if (!error) {
            clearValidationResults();
            clearAuthAndRegData();
            navigate(routes.budget, {replace: true});
        }
    }

    const saveAndAuth = async (event: Event<HTMLFormElement>) => {
        event.preventDefault();

        clearValidationResults();

        validateForm($currentRegData);

        if ($validationResults.length > 0) {
            return;
        }

        const error = await registerUser($currentRegData);

        updateValidationResults(error);

        if (!error) {
            clearValidationResults();
            clearAuthAndRegData();
            navigate(routes.budget, {replace: true});
        }
    }

    const onChangeForm = (isRegistration: boolean) => {
        clearValidationResults();
        isRegistrationForm = isRegistration;
    }

</script>

<form class="{style.Form} {style.FlexVert}">
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
    <div class="{$validationResults.length > 0 ? style.ButtonsBlockWithErrors : style.ButtonsBlock} {style.FlexHor}">
        {#if $validationResults.length > 0}
            <ul class="{style.FlexVert} {style.Font312Red} {style.ValidationBlock}">
                {#each $validationResults as error (`${error.message}`)}
                    <li>{error.message}</li>
                {/each}
            </ul>
        {/if}
        <div class="{style.ButtonsWrapper}">
            {#if !isRegistrationForm}
            <Button title={$_("auth.buttons.register")}
                    onClickHandler={() => onChangeForm(true)}
                    secondary={true}
                            buttonClass={style.AuthButton}/>
            {:else}
            <Button title={$_("auth.buttons.auth")}
                    onClickHandler={() => onChangeForm(false)}
                    secondary={true}
                            buttonClass={style.AuthButton}/>
            {/if}
            <Button title={isRegistrationForm ? $_("auth.buttons.save_and_login") : $_("auth.labels.login")}
                    onClickHandler={isRegistrationForm ? saveAndAuth : onAuth}
                    buttonClass={style.AuthButton}/>
        </div>
    </div>
</form>