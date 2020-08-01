<script lang="ts">
    import {_} from 'svelte-i18n'
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
        ButtonsWrapper
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

    const validateForm = (event: Event<HTMLInputElement>) => {
        console.log("validation")
    }

    const onShopSelect = (selectedId: number) => {
        console.log(selectedId)
    }

    const onAddNewItemToPurchase = () => {
        purchase.goods = [...purchase.goods, {
            category: {}
        }]
    }

    const onRemoveLastItemFromPurchase = () => {
        purchase.goods = [...purchase.goods.slice(0, purchase.goods.length - 1)]
    }

    onMount(() => {
        validateForm(null);
    })

    let purchase: Purchase = {
        shop: {},
        goods: [{
            category: {}
        }],
    };

</script>


<form class="{SideMinorPadding} {Form}">
    <div class="{FlexHor} {MainFieldsWrapper}">
        <div class="{FlexVert}" style="width: 50%; margin-right: 32px">
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
        <div class="{FlexVert}" style="width: 50%">
            <TextAreaWithLabel
                    on:input={validateForm} on:change={validateForm}
                    textAreaClass={TextArea}
                    label={$_("common.labels.comment")} name="comment"
                    bind:value={purchase.comment}/>
        </div>
    </div>
    <SimpleBar style="max-height: {maxHeight}px; width: 100%">
        {#each purchase.goods as goodsItem}
            <div class="{FlexHor} {MinorFieldsWrapper}">
                <div class="{FlexVert}"
                     style="width: 33.333%; margin-right: 32px">
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
                <div class="{FlexVert}"
                     style="width: 33.333%; margin-right: 32px">
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
                <div class="{FlexVert}" style="width: 33.333%">
                    <TextAreaWithLabel
                            on:input={validateForm} on:change={validateForm}
                            textAreaClass={TextArea}
                            label={$_("common.labels.comment")} name="comment"
                            bind:value={goodsItem.comment}/>
                </div>
            </div>
        {/each}
        <div class="{ButtonsWrapper}">
            <ButtonIconNew width={32} height={32}
                           onClickHandler={onAddNewItemToPurchase}/>
            <ButtonIconMinus width={32} height={32}
                             onClickHandler={onRemoveLastItemFromPurchase}/>
        </div>

    </SimpleBar>
</form>