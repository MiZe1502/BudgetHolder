<script lang="ts">
    import {FlexVert, Form, ButtonsBlock, AuthButton} from "./style";

    import InputWithLabel
        from "../../../common/components/Inputs/InputWithLabel/InputWithLabel.svelte";
    import {currentAuthData} from "../../../stores/auth";
    import Button
        from "../../../common/components/Buttons/Button/Button.svelte";

    const onAuth = () => {
        //auth
    }

</script>

<form class="{Form} {FlexVert}">
    <InputWithLabel label="Login" name="login" type="text"
                    bind:value={$currentAuthData.login}/>
    <InputWithLabel label="Password" name="password" type="password"
                    bind:value={$currentAuthData.password}/>

    <div class="{ButtonsBlock}">
        <Button title="Login" onClickHandler={onAuth} buttonClass={AuthButton}/>
    </div>
</form>