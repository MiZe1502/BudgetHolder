<script lang="ts">
    import SimpleBar from '@woden/svelte-simplebar'
    import ClickOutside from "svelte-click-outside";
    import InputWithLabel from "../InputWithLabel/InputWithLabel.svelte";
    import {
        Font312DarkGray,
        Font312Black,
        Input,
        FlexVert,
        Wrapper,
        Font312RedAttention,
        InvalidInput,
        DropdownInput,
        FlexHorCenter,
        ButtonWrapper,
        ArrowIcon,
        ArrowIconExpanded,
        IconButton,
        DropdownInputExpanded
    } from "./style";

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

    const onSelect = (selectedId: number) => {
        console.log(selectedId)
        value = selectedId;
        onSelectHandler && onSelectHandler(selectedId);
        isOpen = false;
    }

    let dropdownElement: HTMLDivElement = null;

    $: inputValue = data.find((elem) => {
        console.log(elem.id, value.id)
        return elem.id === value.id
    }) ? data.find((elem) => elem.id === value.id).value : null;

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

<div class="{FlexVert} {Wrapper}">
    {#if label}
        <label class="{Font312DarkGray} {invalid && Font312RedAttention}"
               for={name}>{label}</label>
    {/if}

    <ClickOutside on:clickoutside={onClickOutside}>
        <div class="{Wrapper} {FlexHorCenter}">
            <input on:input on:change readonly on:click={onClickHandler}
                   class="{Input} {DropdownInput} {isOpen && DropdownInputExpanded} {Font312Black} {invalid && InvalidInput}"
                   {required}
                   {disabled} {name}
                   value={inputValue} {autofocus}/>
            <div class="{ButtonWrapper}">
                <ButtonIconExpandArrow onClickHandler={onClickHandler}
                                       width={16}
                                       height={16}
                                       className="{IconButton} {isOpen ? ArrowIconExpanded : ArrowIcon}"/>
            </div>
        </div>
        {#if isOpen}
            <DropdownMenu value={value} isActiveCondition={(val1, val2) => val1 === val2} data={data} onSelectHandler={onSelect}/>
        {/if}
    </ClickOutside>
</div>