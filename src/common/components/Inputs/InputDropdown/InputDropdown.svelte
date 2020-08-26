<script lang="ts">
    import SimpleBar from '@woden/svelte-simplebar'
    import ClickOutside from "svelte-click-outside";
    import InputWithLabel from "../InputWithLabel/InputWithLabel.svelte";
    import {style} from "./style";

    import {InputDropdownData} from "./utils";
    import ButtonIconExpandArrow
        from "../../Buttons/ButtonIconExpandArrow/ButtonIconExpandArrow.svelte";
    import DropdownMenu from "../../DropdownMenu/DropdownMenu.svelte";

    let isOpen = false;

    const onClickHandler = () => {
        isOpen = !isOpen;
    }

    const onClickOutside = () => {
        isOpen = false;
    }

    const onSelect = (selectedItem) => {
        value = selectedItem.id;
        onSelectHandler && onSelectHandler(selectedItem.id);
        isOpen = false;
    }

    $: inputValue = data.find((elem) => elem.id === value) ? data.find((elem) => elem.id === value).value : null;

    export let data: InputDropdownData[] = [];
    export let onSelectHandler = (selectedId: number) => {
    };
    export let value: number = -1;
    export let autofocus = false;
    export let name = "";
    export let label = "";
    export let disabled = false;
    export let required = false;
    export let invalid = false;
</script>

<div class="{style.FlexVert} {style.Wrapper}">
    {#if label}
        <label class="{style.Font312DarkGray} {invalid && style.Font312RedAttention}"
               for={name}>{label}</label>
    {/if}

    <ClickOutside on:clickoutside={onClickOutside}>
        <div class="{style.Wrapper} {style.FlexHorCenter}">
            <input on:input on:change readonly on:click={onClickHandler}
                   class="{style.Input} {style.DropdownInput} {isOpen && style.DropdownInputExpanded} {style.Font312Black} {invalid && style.InvalidInput}"
                   {required}
                   {disabled} {name}
                   value={inputValue} {autofocus}/>
            <div class="{style.ButtonWrapper}">
                <ButtonIconExpandArrow onClickHandler={onClickHandler}
                                       width={16}
                                       height={16}
                                       className="{style.IconButton} {isOpen ? style.ArrowIconExpanded : style.ArrowIcon}"/>
            </div>
        </div>
        {#if isOpen}
            <DropdownMenu value={value}
                          isActiveCondition={(val1, val2) => val1 === val2}
                          data={data} onSelectHandler={onSelect}/>
        {/if}
    </ClickOutside>
</div>