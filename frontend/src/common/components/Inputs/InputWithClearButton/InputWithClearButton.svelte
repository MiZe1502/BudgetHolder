<script lang="ts">
    import InputWithIcon from "../InputWithIcon/InputWithIcon.svelte";
    import ButtonIconClear
        from "../../Buttons/ButtonIconClear/ButtonIconClear.svelte";

    const clearValue = () => {
        value = "";
        if (required) {
            invalid = true;
        }
    }

    export let value = "";
    export let autofocus = false;
    export let name = "";
    export let type = "text";
    export let label = "";
    export let disabled = false;
    export let required = false;
    export let invalid = false;
</script>

<InputWithIcon on:input on:change bind:value={value}
               {autofocus} {name} {type} {invalid}
               {label} {disabled} {required}>
    <ButtonIconClear onClickHandler={clearValue}/>
</InputWithIcon>