<script lang="ts">
    import SimpleBar from '@woden/svelte-simplebar'
    import InputWithLabel from "../InputWithLabel/InputWithLabel.svelte";
    import {
        Font312DarkGray,
        Font312Black,
        Input,
        FlexVert,
        Wrapper,
        Font312RedAttention,
        InvalidInput,
        Dropdown,
        DropdownInput,
        SingleLine,
        FlexHorCenter,
        maxHeight,
        SelectedLine,
        ButtonWrapper,
        ArrowIcon,
        ArrowIconExpanded,
        IconButton
    } from "./style";

    import {InputDropdownData} from "./utils";
    import ButtonIconExpandArrow
        from "../../Buttons/ButtonIconExpandArrow/ButtonIconExpandArrow.svelte";

    let isOpen = false;

    const onClickHandler = () => {
        isOpen = !isOpen;
    }

    const onSelect = (selectedId: number) => {
        value = selectedId;
        onSelectHandler && onSelectHandler(selectedId);
        isOpen = false;
    }

    export let data: InputDropdownData[] = [];
    export let onSelectHandler = (selectedId: number) => {
    };
    export let value: number = 0;
    export let autofocus = false;
    export let name = "";
    export let label = "";
    export let disabled = false;
    export let required = false;
    export let invalid = false;</script>

<div class="{FlexVert} {Wrapper}">
    {#if label}
        <label class="{Font312DarkGray} {invalid && Font312RedAttention}"
               for={name}>{label}</label>
    {/if}

    <div class="{Wrapper} {FlexHorCenter}">
        <input on:input on:change readonly on:click={onClickHandler}
               class="{Input} {DropdownInput} {Font312Black} {invalid && InvalidInput}"
               {required}
               {disabled} {name}
               value={data && data.find((elem) => elem.id === value).value} {autofocus}/>
        <div class="{ButtonWrapper}">
            <ButtonIconExpandArrow onClickHandler={onClickHandler} width={16}
                                   height={16}
                                   className="{IconButton} {isOpen ? ArrowIconExpanded : ArrowIcon}"/>
        </div>
    </div>
    {#if isOpen}
        <div class={Dropdown}>
            <SimpleBar style="max-height: {maxHeight}px; width: 100%">
                {#each data as item (item.id)}
                    <div on:click={() => onSelect(item.id)}
                         class="{item.id === value && SelectedLine} {SingleLine} {FlexHorCenter} {Font312Black}">{item.value}</div>
                {/each}
            </SimpleBar>
        </div>
    {/if}
</div>