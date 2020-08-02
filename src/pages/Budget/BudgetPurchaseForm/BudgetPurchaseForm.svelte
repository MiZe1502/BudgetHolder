<script lang="ts">
    import {_} from 'svelte-i18n'
    import {v4 as uuidv4} from 'uuid';
    import SimpleBar from '@woden/svelte-simplebar'
    import {onMount} from "svelte";
    import {Purchase} from "../types";
    import {getSimpleShopsData} from "../../../stores/shops";
    import {
        SideMinorPadding,
        TextArea,
        FlexVert,
        FlexHor,
        Form,
        MainFieldsWrapper,
        MinorFieldsWrapper,
        maxHeight,
        ButtonsWrapper,
        MainColumn,
        MinorColumn,
        NotLastColumn
    } from "./style";

    import {simpleCategories} from "../../../stores/categories";
    import InputWithLabel
        from "../../../common/components/Inputs/InputWithLabel/InputWithLabel.svelte";
    import TextAreaWithLabel
        from "../../../common/components/Inputs/TextAreaWithLabel/TextAreaWithLabel.svelte";
    import InputDropdown
        from "../../../common/components/Inputs/InputDropdown/InputDropdown.svelte";
    import ButtonIconNew
        from "../../../common/components/Buttons/ButtonIconNew/ButtonIconNew.svelte";
    import ButtonIconMinus
        from "../../../common/components/Buttons/ButtonIconMinus/ButtonIconMinus.svelte";
    import Button
        from "../../../common/components/Buttons/Button/Button.svelte";

    const validateForm = (event: Event<HTMLInputElement>) => {
        console.log("validation")
    }

    const onShopSelect = (selectedId: number) => {
        console.log(selectedId)
    }

    const onAddNewItemToPurchase = () => {
        purchase.goods = [...purchase.goods, {
            tempId: uuidv4(),
            category: {}
        }]
    }

    const onRemoveItemFromPurchase = (tempId: string) => {
        purchase.goods = [...purchase.goods.filter((goodsItem) => goodsItem.tempId !== tempId)]
    }

    const onClearForm = () => {
        purchase = Object.assign({}, {
            shop: {},
            goods: [{
                tempId: uuidv4(),
                category: {}
            }],
        })
    }

    onMount(() => {
        validateForm(null);
    })

    let purchase: Purchase = {
        shop: {},
        goods: [{
            tempId: uuidv4(),
            category: {}
        }],
    };

</script>


<form class="{Form}">
    <div class="{SideMinorPadding} {FlexHor} {MainFieldsWrapper}">
        <div class="{FlexVert} {MainColumn} {NotLastColumn}">
            <InputWithLabel
                    on:input={validateForm} on:change={validateForm}
                    label={$_("budget.labels.price")} autofocus={true}
                    type="number" name="totalPrice"
                    bind:value={purchase.totalPrice} required={true}/>
            <InputWithLabel
                    on:input={validateForm} on:change={validateForm}
                    label={$_("budget.labels.date")} autofocus={true}
                    type="date" name="date"
                    bind:value={purchase.date} required={true}/>
            <InputDropdown onSelectHandler={onShopSelect}
                           bind:value={purchase.shop.id} name="shop"
                           label={$_("budget.labels.shop")}
                           data={getSimpleShopsData()}/>
        </div>
        <div class="{FlexVert} {MainColumn}">
            <TextAreaWithLabel
                    on:input={validateForm} on:change={validateForm}
                    textAreaClass={TextArea}
                    label={$_("common.labels.comment")} name="comment"
                    bind:value={purchase.comment}/>
        </div>
    </div>
    <SimpleBar style="max-height: {maxHeight}px; width: 100%">
        {#each purchase.goods as goodsItem (goodsItem.tempId)}
            <div class="{FlexVert} {MinorFieldsWrapper}">
                <div class="{FlexHor}">
                    <div class="{FlexVert} {MinorColumn} {NotLastColumn}">
                        <InputWithLabel
                                on:input={validateForm} on:change={validateForm}
                                label={$_("common.labels.name")}
                                type="text" name="name"
                                bind:value={goodsItem.name}/>
                        <InputDropdown onSelectHandler={onShopSelect}
                                       bind:value={goodsItem.category.id}
                                       name="category"
                                       label={$_("budget.labels.category")}
                                       data={$simpleCategories}/>
                    </div>
                    <div class="{FlexVert} {MinorColumn} {NotLastColumn}">
                        <InputWithLabel
                                on:input={validateForm} on:change={validateForm}
                                label={$_("budget.labels.amount")}
                                type="number" name="amount"
                                bind:value={goodsItem.amount}/>
                        <InputWithLabel
                                on:input={validateForm} on:change={validateForm}
                                label={$_("budget.labels.single_price")}
                                type="number" name="price"
                                bind:value={goodsItem.price}/>
                    </div>
                    <div class="{FlexVert} {MinorColumn}">
                        <TextAreaWithLabel
                                on:input={validateForm} on:change={validateForm}
                                textAreaClass={TextArea}
                                label={$_("common.labels.comment")}
                                name="comment"
                                bind:value={goodsItem.comment}/>
                    </div>
                </div>
                <div class="{FlexHor}">
                    <ButtonIconMinus width={32} height={32}
                                     onClickHandler={() => onRemoveItemFromPurchase(goodsItem.tempId)}/>
                </div>
            </div>
        {/each}
    </SimpleBar>
    <div class="{ButtonsWrapper} {FlexHor}">
        <ButtonIconNew width={32} height={32}
                       onClickHandler={onAddNewItemToPurchase}/>
        <div>
            <Button title={$_("budget.buttons.save")}/>
            <Button onClickHandler={onClearForm} secondary={true}
                    title={$_("budget.buttons.clear")}/>
        </div>
    </div>

</form>