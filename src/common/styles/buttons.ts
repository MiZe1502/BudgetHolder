import { css } from 'emotion'

import colors from "./colors";

export const defaultButtonIconSize = 24;

export const ButtonReset = css`
    border: none;
    background: transparent;
    text-decoration: none;
    outline: none;
    padding: 0;
`;

export const ButtonNegative = css`
    background-color: ${colors.redAttention}; 
    fill: ${colors.redAttention}; 
    
    &:hover,
    &:active {
        background-color: ${colors.redActive}; 
        fill: ${colors.redActive}; 
    }
    
    &:active {
        background-color: ${colors.redLight};
        fill: ${colors.redLight};
    }
`

export const ButtonSecondary = css`
    background-color: ${colors.orangePrimary}; 
    fill: ${colors.orangePrimary}; 
    
    &:hover,
    &:active {
        background-color: ${colors.orangeActive}; 
        fill: ${colors.orangeActive}; 
    }
    
    &:active {
        background-color: ${colors.orangeLight};
        fill: ${colors.orangeLight};
    }
`;

export const ButtonPrimary = css`
	background-color: ${colors.greenPrimary}; 
	fill: ${colors.greenPrimary}; 
	
	&:hover,
    &:active {
        background-color: ${colors.greenActive}; 
        fill: ${colors.greenActive}; 
    }
    
    &:active {
        background-color: ${colors.greenLight};
        fill: ${colors.greenLight};
    }
`;

export const ButtonDefault = css`
	transition: background-color 300ms ease;
    cursor: pointer;

	&:hover,
	&:active { 
        cursor: pointer;
	}

    &:disabled {
		background-color: ${colors.darkGray};
		fill: ${colors.darkGray};
        border: none;
        cursor: default;
        box-shadow: none;
        color: whitesmoke;
	}
`;