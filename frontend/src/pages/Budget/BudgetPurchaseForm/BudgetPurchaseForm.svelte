<script lang="ts">
    import {_} from 'svelte-i18n'
    import {v4 as uuidv4} from 'uuid';
    import SimpleBar from '@woden/svelte-simplebar'
    import {onMount} from "svelte";
    import {Purchase} from "../types";
    import {SuggestionItem} from "../../../common/components/Inputs/InputWithSuggestion/utils";
    import {getShopById, getSimpleShopsData} from "../../Shops/shops";
    import {style} from "./style";

    import {
        getSimpleCategoryById,
        simpleCategories
    } from "../../Categorization/categories";
    import {simpleShops} from "../../Shops/shops";
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
    import {
        addPurchaseToStore,
        currentPurchase,
        clearCurrentPurchaseData,
        purchaseLocalStorageKey,
        purchaseLocalStorageUpdateInterval,
        updateValidationResults,
        validationResults,
        clearValidationResults
    } from "../../../stores/purchases";
    import {
        goodsForSuggestions,
        goods,
        addGoodsToStore
    } from "../../../stores/goods";
    import {
        addDataToLocalStorage,
        getDataFromLocalStorageByKey
    } from "../../../common/utils/localStorage";
    import InputWithSuggestion
        from "../../../common/components/Inputs/InputWithSuggestion/InputWithSuggestion.svelte";
    import {
        validationRulesGoods,
        validationRulesPurchase,
        arePricesFit
    } from "./validationRules";

    const validateForm = () => {
        clearValidationResults();

        let pricesFit = arePricesFit($currentPurchase);

        if (!pricesFit) {
            updateValidationResults("Total purchase price not fits to summary price of all goods")
        }

        for (let rule of validationRulesPurchase) {
            const isInvalid = rule.validator($currentPurchase);
            if (isInvalid) {
                updateValidationResults(rule.message)
            }
        }

        for (let i = 0; i < $currentPurchase.goods_data.length; i++) {
            for (let rule of validationRulesGoods) {
                const isInvalid = rule.validator($currentPurchase.goods_data[i]);
                if (isInvalid) {
                    updateValidationResults(rule.message, i)
                }
            }
        }
    }

    const onGoodsItemSelect = (selectedItem: SuggestionItem, goodsItemTempId: string) => {
        let goodsItemDetails = $currentPurchase.goods_data.find((item) => item.tempId === goodsItemTempId);
        const goodsItemData = $goods.find((item) => item.id === selectedItem.id);

        if (goodsItemData && goodsItemDetails) {
            goodsItemDetails = Object.assign(goodsItemDetails, goodsItemData);
        }
    }

    const onShopSelect = async (selectedId: number) => {
        const shop = await getShopById(selectedId)
        $currentPurchase.shop = {...shop};
        $currentPurchase.shop_id = shop.id;
    }

    const onAddNewItemToPurchase = (event: MouseEvent) => {
        event.preventDefault();

        $currentPurchase.goods = [...$currentPurchase.goods_data, {
            tempId: uuidv4(),
            category: {}
        }]
    }

    const onRemoveItemFromPurchase = (tempId: string) => {
        $currentPurchase.goods = [...$currentPurchase.goods_data.filter((goodsItem) => goodsItem.tempId !== tempId)]
    }

    const onCategorySelect = (goodsItemTempId: number, selectedId: number) => {
        const updatedGoodsItem = purchase.goods_data.find((goodsItem) => goodsItem.tempId === goodsItemTempId);
        updatedGoodsItem.category = getSimpleCategoryById(selectedId);
    }

    const onClearForm = (event: MouseEvent) => {
        event.preventDefault();

        clearCurrentPurchaseData();
        clearValidationResults();
    }

    const onSave = async (event: MouseEvent) => {
        event.preventDefault();

        validateForm();

        if ($validationResults.length > 0) {
            return;
        }

        await addPurchaseToStore(purchase);
        //addGoodsToStore(purchase.goods);
        clearValidationResults();
    }

    onMount(() => {
        if (!getDataFromLocalStorageByKey(purchaseLocalStorageKey)) {
            addDataToLocalStorage(purchaseLocalStorageKey, $currentPurchase);
        }

        setInterval(() => {
            addDataToLocalStorage(purchaseLocalStorageKey, $currentPurchase);
        }, purchaseLocalStorageUpdateInterval)
    })

    let purchase: Purchase = $currentPurchase;

</script>


<form class="{style.Form}">
    <div class="{style.SideMinorPadding} {style.FlexHor} {style.MainFieldsWrapper}">
        <div class="{style.FlexVert} {style.MainColumn} {style.NotLastColumn}">
            <InputWithLabel
                    label={$_("budget.labels.price")} autofocus={true}
                    type="number" name="total_price"
                    bind:value={$currentPurchase.total_price} required={true}/>
            <InputWithLabel
                    label={$_("budget.labels.date")} autofocus={true}
                    type="date" name="date"
                    bind:value={$currentPurchase.date} required={true}/>
            <InputDropdown onSelectHandler={onShopSelect}
                           bind:value={$currentPurchase.shop.id} name="shop"
                           label={$_("budget.labels.shop")}
                           data={$simpleShops}/>
        </div>
        <div class="{style.FlexVert} {style.MainColumn}">
            <TextAreaWithLabel
                    textAreaClass={style.TextArea}
                    label={$_("common.labels.comment")} name="comment"
                    bind:value={$currentPurchase.comment}/>
        </div>
    </div>
    <SimpleBar style="max-height: {style.maxHeight}px; width: 100%">
        {#each $currentPurchase.goods_data as goodsItem, index (goodsItem.tempId)}
            <div class="{style.FlexVert} {style.MinorFieldsWrapper}">
                <div class="{style.FlexHor}">
                    <div class="{style.FlexVert} {style.MinorColumn} {style.NotLastColumn}">
                        <InputWithSuggestion
                                onSelectHandler={(selectedItem) => onGoodsItemSelect(selectedItem, goodsItem.tempId)}
                                suggestionsList={$goodsForSuggestions}
                                        label={$_("common.labels.name")}
                                        type="text" name="name"
                                bind:value={goodsItem.name}/>
                        <InputDropdown
                                onSelectHandler={(selectedId) => onCategorySelect(goodsItem.tempId, selectedId)}
                                bind:value={goodsItem.category.id}
                                        name="category"
                                label={$_("budget.labels.category")}
                                        data={$simpleCategories}/>
                    </div>
                    <div class="{style.FlexVert} {style.MinorColumn} {style.NotLastColumn}">
                        <InputWithLabel
                                label={$_("budget.labels.amount")}
                                type="number" name="amount"
                                bind:value={goodsItem.amount}/>
                        <InputWithLabel
                                label={$_("budget.labels.single_price")}
                                type="number" name="price"
                                bind:value={goodsItem.price}/>
                    </div>
                    <div class="{style.FlexVert} {style.MinorColumn}">
                        <TextAreaWithLabel
                                textAreaClass={style.TextArea}
                                label={$_("common.labels.comment")}
                                name="comment"
                                bind:value={goodsItem.comment}/>
                    </div>
                </div>
                <div class="{style.FlexHorCenter} {style.GoodsItemControlBlock}">
                    <ButtonIconMinus width={24} height={24}
                                     onClickHandler={() => onRemoveItemFromPurchase(goodsItem.tempId)}/>
                    <div class="{style.FlexHorCenter} {style.CounterBlock} {style.Font312Black}">
                        {index + 1}
                    </div>
                </div>
            </div>
        {/each}
    </SimpleBar>
    <div class="{style.ButtonsWrapper} {style.FlexHor}">
        <ButtonIconNew width={24} height={24}
                       onClickHandler={onAddNewItemToPurchase}/>
        {#if $validationResults.length > 0}
            <ul class="{style.FlexVert} {style.Font312Red} {style.ValidationBlock}">
                {#each $validationResults as error (`${error.goodsItemCounter}-${error.message}`)}
                    <li>{error.goodsItemCounter ? `Item № ${error.goodsItemCounter}: `: ""}{error.message}</li>
                {/each}
            </ul>
        {/if}
        <div class="{style.FlexHorCenter} {style.ButtonsBlock}">
            <Button disabled={!$currentPurchase.total_price}
                    buttonClass="{style.ButtonForm}"
                    title={$_("budget.buttons.save")} onClickHandler={onSave}/>
            <Button onClickHandler={onClearForm} buttonClass={style.ButtonForm}
                    secondary={true}
                    title={$_("budget.buttons.clear")}/>
        </div>
    </div>

</form>